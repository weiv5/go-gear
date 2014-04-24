package gear

import (
    "net/http"
)

type AppInterface interface {
    Init(http.ResponseWriter, *Request)
    IndexAction()
}

type App struct {
    Request
    Response
}

func (app *App) Init(w http.ResponseWriter, r *Request) {
    app.W = w
    app.R = r.R
    app.Module = r.Module
    app.Action = r.Action
    app.Data = make(map[string] interface{})
}
