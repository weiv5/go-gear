package app

import (
    "gear"
)

type TestApp struct {
    gear.App
    CheckIp
}


func (app *TestApp) IndexAction() {
    tplPath := "src/web/view/"
    app.Data["ip"] = app.Ip2Long()
    app.Display("content", tplPath+"header.html", tplPath+"footer.html", tplPath+"content.html")
}


type Ret struct {
    Status   bool   `json:"status"`
    Data     interface{} `json:"data"`
}

type DataM struct {
    Ip string
    What string
}

func (app *TestApp) GoAction() {
    what := app.Get("what")
    ret := &Ret{false, &DataM{app.Ip(), what}}
    app.Json(ret)
}
