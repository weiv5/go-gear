package app

import (
    "gear"
)

type TestApp struct {
    gear.App
}

var (
    tplPath = "src/web/view/"
)

func (app *TestApp) IndexAction() {
    app.Assign("ip", app.Ip2Long())
    app.Display("content", tplPath+"header.html", tplPath+"footer.html", tplPath+"content.html")
}





type Ret struct {
    Status   bool   `json:"status"`
    Data     string `json:"data"`
}

func (app *TestApp) GoAction() {
    ret := &Ret{false, app.Ip()}
    app.Json(ret)
}
