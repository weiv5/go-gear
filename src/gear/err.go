package gear

import (
    "os"
    "runtime"
    "path/filepath"
    "strings"
    "net/http"
    "fmt"
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

func (errM *ErrM) NotFound(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(404)
    fmt.Fprintln(w, "404 page not found")
    return
}

var (
    Err = &ErrM{}
)
