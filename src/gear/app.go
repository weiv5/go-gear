package gear

import (
    "gear/session"
)

type AppInterface interface {
    Init(*Response, *Request)
}

type App struct {
    Request
    Response
    Session session.SessionStore
}

func (app *App) Init(w *Response, r *Request) {
    app.W = w.W
    app.R = r.R
    app.Module = r.Module
    app.Action = r.Action
    app.Data = make(map[string] interface{})
    app.Session = r.Session
}
