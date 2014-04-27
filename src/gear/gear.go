package gear

import (
    "net/http"
    "time"
    "runtime"
    "gear/session"
)

var (
    SessionOn bool = false
    SessionM *session.Manager
)

func SessionStart() {
    SessionOn = true
    cookie := Ini.String("session::cookiename")
    gclifetime := Ini.String("session::gclifetime")
    provider := Ini.String("session::savepath")
    conf := `{"cookieName":"`+cookie+`","gclifetime":`+gclifetime+`,"providerConfig":"`+provider+`"}`
    var err error
    SessionM, err = session.NewManager("file", conf)
    if err != nil {
        Log.WriteLog(err)
        go SessionM.GC()
    }
}


func Run() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    Log.WatchPanic()
    go Log.WatchAccess()

    r_timeout := Ini.Int("r_timeout")
    w_timeout := Ini.Int("w_timeout")
    s := &http.Server{
        Addr        : Ini.String("addr"),
        Handler     : &Serve{},
        ReadTimeout : time.Duration(r_timeout) * time.Second,
        WriteTimeout: time.Duration(w_timeout) * time.Second,
    }
    err := s.ListenAndServe()
    if err != nil {
        Log.WriteLog(err)
    }
}
