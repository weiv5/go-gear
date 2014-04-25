package gear

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
    "flag"
)

var (
	DEFAULT_SECTION = "default"   // default section means if some ini items not in a section, make them in default section,
	bNumComment     = []byte{'#'} // number signal
	bSemComment     = []byte{';'} // semicolon signal
	bEmpty          = []byte{}
	bEqual          = []byte{'='} // equal signal
	bDQuote         = []byte{'"'} // quote signal
	sectionStart    = []byte{'['} // section start signal
	sectionEnd      = []byte{']'} // section end signal
)

// IniConfig implements Config to parse ini file.
type IniConfig struct {}

// ParseFile creates a new Config and parses the file configuration from the named file.
func (ini *IniConfig) Parse(name string) (*IniConfigContainer, error) {
    fileInfo, err := os.Stat(name)
    if err != nil {
        return nil, err
    }
    if fileInfo.IsDir() {
        return nil, errors.New("ini file can not be dir")
    }
    file, err := os.Open(name)
    if err != nil {
        return nil, err
    }

    cfg := &IniConfigContainer{
        file.Name(),
        make(map[string]map[string]string),
        make(map[string]string),
        make(map[string]string),
        sync.RWMutex{},
    }
    cfg.Lock()
    defer cfg.Unlock()
    defer file.Close()

    var comment bytes.Buffer
    buf := bufio.NewReader(file)
    section := DEFAULT_SECTION
    for {
        line, _, err := buf.ReadLine()
        if err == io.EOF {
            break
        }
        if bytes.Equal(line, bEmpty) {
            continue
        }
        line = bytes.TrimSpace(line)

        var bComment []byte
        switch {
        case bytes.HasPrefix(line, bNumComment):
            bComment = bNumComment
        case bytes.HasPrefix(line, bSemComment):
            bComment = bSemComment
        }
        if bComment != nil {
            line = bytes.TrimLeft(line, string(bComment))
            line = bytes.TrimLeftFunc(line, unicode.IsSpace)
            comment.Write(line)
            comment.WriteByte('\n')
            continue
        }

        if bytes.HasPrefix(line, sectionStart) && bytes.HasSuffix(line, sectionEnd) {
            section = string(line[1 : len(line)-1])
            section = strings.ToLower(section) // section name case insensitive
            if comment.Len() > 0 {
                cfg.sectionComment[section] = comment.String()
                comment.Reset()
            }
            if _, ok := cfg.data[section]; !ok {
                cfg.data[section] = make(map[string]string)
            }
        } else {
            if _, ok := cfg.data[section]; !ok {
                cfg.data[section] = make(map[string]string)
            }
            keyval := bytes.SplitN(line, bEqual, 2)
            val := bytes.TrimSpace(keyval[1])
            if bytes.HasPrefix(val, bDQuote) {
                val = bytes.Trim(val, `"`)
            }

            key := string(bytes.TrimSpace(keyval[0])) // key name case insensitive
            key = strings.ToLower(key)
            cfg.data[section][key] = string(val)
            if comment.Len() > 0 {
                cfg.keycomment[section+"."+key] = comment.String()
                comment.Reset()
            }
        }

    }
    return cfg, nil
}

// A Config represents the ini configuration.
// When set and get value, support key as section:name type.
type IniConfigContainer struct {
    filename       string
    data           map[string]map[string]string // section=> key:val
    sectionComment map[string]string            // section : comment
    keycomment     map[string]string            // id: []{comment, key...}; id 1 is for main comment.
    sync.RWMutex
}

// Bool returns the boolean value for a given key.
func (c *IniConfigContainer) Bool(key string) bool {
    key = strings.ToLower(key)
    b, err := strconv.ParseBool(c.getdata(key))
    if err != nil {
        return false
    }
    return b
}

// Int returns the integer value for a given key.
func (c *IniConfigContainer) Int(key string) int {
    key = strings.ToLower(key)
    i, err := strconv.Atoi(c.getdata(key))
    if err != nil {
        return 0
    }
    return i
}

// Int64 returns the int64 value for a given key.
func (c *IniConfigContainer) Int64(key string) int64 {
    key = strings.ToLower(key)
    i, err := strconv.ParseInt(c.getdata(key), 10, 64)
    if err != nil {
        return 0
    }
    return i
}

// Float returns the float value for a given key.
func (c *IniConfigContainer) Float(key string) float64 {
    key = strings.ToLower(key)
    f,err := strconv.ParseFloat(c.getdata(key), 64)
    if err != nil {
        return 0
    }
    return f
}

// String returns the string value for a given key.
func (c *IniConfigContainer) String(key string) string {
    key = strings.ToLower(key)
    return c.getdata(key)
}

// Strings returns the []string value for a given key.
func (c *IniConfigContainer) Strings(key string) []string {
    return strings.Split(c.String(key), ";")
}


// section.key or key
func (c *IniConfigContainer) getdata(key string) string {
    c.RLock()
    defer c.RUnlock()
    if len(key) == 0 {
        return ""
    }

    var section, k string
    key = strings.ToLower(key)
    sectionkey := strings.Split(key, "::")
    if len(sectionkey) >= 2 {
        section = sectionkey[0]
        k = sectionkey[1]
    } else {
        section = DEFAULT_SECTION
        k = sectionkey[0]
    }
    if v, ok := c.data[section]; ok {
        if vv, o := v[k]; o {
            return vv
        }
    }
    return ""
}

var (
    IniFile = flag.String("conf", "", "ini file path")
    IniParse = &IniConfig{}
    Ini *IniConfigContainer
)

func init() {
    var err error
    flag.Parse()
    Ini,err = IniParse.Parse(*IniFile)
    if err != nil {
        Log.WriteLog(err)
        os.Exit(1)
    }
}
