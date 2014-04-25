package app

import (
    "gear"
)

type LoginApp struct {
    gear.App
}


func (app *LoginApp) IndexAction() {
    tplPath := "src/web/view/"
    app.Display("content", tplPath+"header.html", tplPath+"footer.html", tplPath+"login.html")
}
