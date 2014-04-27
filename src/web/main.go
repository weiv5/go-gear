package main

import (
    "gear"
    "web/app"
)

func main() {
    gear.StaticRoute()
    gear.AddRoute("/", &app.IndexApp{})
    gear.AddRoute("/login", &app.LoginApp{})
    gear.AddRoute("/info", &app.InfoApp{})
    gear.SessionStart()
    gear.Run()
}
