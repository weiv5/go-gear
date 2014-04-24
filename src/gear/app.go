package gear

type AppInterface interface {
    Init(*Response, *Request)
    IndexAction()
}

type App struct {
    Request
    Response
}

func (app *App) Init(w *Response, r *Request) {
    app.W = w.W
    app.R = r.R
    app.Module = r.Module
    app.Action = r.Action
    app.Data = make(map[string] interface{})
}
