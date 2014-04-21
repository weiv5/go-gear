package gear

import (
    "os"
    "strconv"
    "syscall"
)

type LogM struct {}

func (l *LogM) WritePid() (err error) {
    pid := strconv.Itoa(os.Getpid())
    if pid == "" {
        return
    }
    filename := "log/gear.pid"
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        return
    }
    defer func() {
        file.Close()
    }()
    file.WriteString(pid)

    return
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
}

var (
    Log = &LogM{}
)

func init() {
    Log.WritePid()
    Log.WatchPanic()
}
