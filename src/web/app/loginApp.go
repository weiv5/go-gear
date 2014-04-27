package app

import (
    "gear"
    "fmt"
)

type LoginApp struct {
    gear.App
}


func (app *LoginApp) IndexAction() {
    tplPath := "src/web/view/"
    app.Display("content", tplPath+"header.html", tplPath+"footer.html", tplPath+"login.html")
}

func (app *LoginApp) VerifyAction() {
    app.Session.Set("uid", 1)
    app.Session.Set("name", app.Post("name"))
    fmt.Println(app.Post("name"), app.Session.Get("uid"))
    app.Redirect("/info")
}
