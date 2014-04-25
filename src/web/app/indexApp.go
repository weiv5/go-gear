package app

import (
    "gear"
)

type IndexApp struct {
    gear.App
}


func (app *IndexApp) IndexAction() {
    tplPath := "src/web/view/"
    app.Data["hello"] = "i'm the index request by"+app.Ip()+" at "+gear.Date("y-m-d")
    app.Display("index", tplPath+"header.html", tplPath+"footer.html", tplPath+"index.html")
}
