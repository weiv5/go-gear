package gear

import (
    "time"
)

func Date(f ...string) string {
    var format string
    if len(f)==1 {
        for _, v := range f[0] {
            switch (v) {
                case 'Y' :
                case 'y' : format += "2006"
                case 'm' : format += "01"
                case 'd' : format += "02"
                case 'H' :
                case 'h' : format += "15"
                case 'i' : format += "04"
                case 's' : format += "05"
                default : format += string(v)
            }
        }
    } else {
        format = "2006-01-02 15:04:05"
    }
    return time.Now().Format(format)
}

func Time() int64 {
    return time.Now().Unix();
}
