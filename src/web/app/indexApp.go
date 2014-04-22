package app

import (
    "gear"
)

type IndexApp struct {
    gear.App
}


func (app *IndexApp) IndexAction() {
    tplPath := "src/web/view/"
    app.Data["ip"] = "i'm the index request by"+app.Ip()
    app.Display("content", tplPath+"header.html", tplPath+"footer.html", tplPath+"content.html")
}
