package app

import (
    "gear"
    "fmt"
)

type TestApp struct {
    gear.App
}

func (app *TestApp) IndexAction() {
    fmt.Println("testApp Index By wei")
    fmt.Println(app.Get("what"))
}

func (app *TestApp) GoAction() {
    fmt.Println("testApp Go By wei")
    fmt.Println(app.Get("what"))
}
