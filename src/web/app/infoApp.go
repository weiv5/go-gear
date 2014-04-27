package app

import (
    "gear"
)

type InfoApp struct {
    gear.App
    CheckLogin
}

func (app *InfoApp) IndexAction() {
    app.Session.Set("uid", 1)
    tplPath := "src/web/view/"
    app.Data["login"] = app.Get("login")
    app.Data["name"] = app.Post("name")
    app.Display("info", tplPath+"header.html", tplPath+"footer.html", tplPath+"info.html")
}

func (app *InfoApp) DetailAction() {
    tplPath := "src/web/view/"
    app.Data["name"] = app.Get("name")
    app.Data["ip"] = app.Ip()
    app.Data["ua"] = app.GetHeader("User-Agent")
    app.Display("detail", tplPath+"header.html", tplPath+"footer.html", tplPath+"detail.html")
}
