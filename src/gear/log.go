package gear

import (
    "os"
    "syscall"
    "runtime"
    "path/filepath"
    "strings"
    "fmt"
    "time"
)

type LogM struct {
    accessChan chan string
}

func (l *LogM) WatchPanic() {
    var panicFile *os.File
    panicFile, err := os.OpenFile("log/panic", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        panicFile, err = os.OpenFile("/dev/null", os.O_RDWR, 0)
    }
    if err == nil {
        fd := panicFile.Fd()
        syscall.Dup2(int(fd), int(os.Stderr.Fd()))
    }
    defer panicFile.Close()
}

func (l *LogM) WriteLog(err error) {
    _, file, line, _ := runtime.Caller(1)
    file = filepath.Base(file)
    errMsg := err.Error()
    if !strings.HasPrefix(errMsg, "DEBUG") && !strings.HasPrefix(errMsg, "INFO") {
        errMsg = "ERROR " + errMsg
    }
    msg := fmt.Sprintf("%s [%s:%d]: %s\n", Date(), file, line, errMsg)
    logFile, err := os.OpenFile("log/log", os.O_CREATE|os.O_WRONLY, 0644)
    defer logFile.Close()
    if err != nil {
        return
    }
    logFile.WriteString(msg)
}

func (l *LogM) WatchAccess() {
    now := time.Now()
    y, m, d := now.Year(), now.Month(), now.Day()
    logFile, err := os.OpenFile("log/access."+fmt.Sprintf("%d%02d%02d", y, m, d), os.O_CREATE|os.O_WRONLY, 0644)
    defer logFile.Close()
    if err != nil {
        return
    }
    for msg := range l.accessChan {
        now2 := time.Now()
        y2, m2, d2 := now2.Year(), now2.Month(), now2.Day()
        if y!=y2 || m!=m2 || d!=d2 {
            y, m, d = y2, m2, d2
            logFile,_ = os.OpenFile("log/access."+fmt.Sprintf("%d%02d%02d", y, m, d), os.O_CREATE|os.O_WRONLY, 0644)
        }
        logFile.WriteString(msg)
    }
}

func (l *LogM) Access(r *Request) {
    msg := fmt.Sprintf("%s\t\"%s\"\t\"%s\"\t\"%s\"\n", Date(), r.Ip(), r.Url(), r.Header("User-Agent"))
    l.accessChan <- msg
}

var (
    Log = &LogM{accessChan:make(chan string, 1024)}
)
