package gear

import (
    "os"
    "fmt"
    "runtime"
    "path/filepath"
    "strings"
)

type ErrM struct {}

func (errM *ErrM) Fatal(err error) {
    _, file, line, _ := runtime.Caller(1)
    file = filepath.Base(file)
    errMsg := err.Error()
    if !strings.HasPrefix(errMsg, "DEBUG") && !strings.HasPrefix(errMsg, "INFO") {
        errMsg = "ERROR " + errMsg
    }
    fmt.Printf("%s [%s:%d]: %s\n", Func.Date(), file, line, errMsg)
    os.Exit(1)
}

var (
    Err = &ErrM{}
)
