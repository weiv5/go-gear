package gear

import (
    "net/http"
    "time"
)

func Run() {
    Log.WatchPanic()
    go Log.WatchAccess()

    r_timeout,_ := Ini.Int("r_timeout")
    w_timeout,_ := Ini.Int("w_timeout")
    s := &http.Server{
        Addr        : Ini.String("addr"),
        Handler     : &Serve{},
        ReadTimeout : time.Duration(r_timeout) * time.Second,
        WriteTimeout: time.Duration(w_timeout) * time.Second,
    }
    s.ListenAndServe()
}
