package gear

import (
    "net/http"
)

type AppInterface interface {
    Init(http.ResponseWriter, *http.Request, string, string)
    IndexAction()
}

type App struct {
    Module string
    Action string
    Request
    Response
}

func (app *App) Init(w http.ResponseWriter, r *http.Request, module string, action string) {
    app.W = w
    app.R = r
    app.Data = make(map[string] interface{})
    app.Module = module
    app.Action = action
}
