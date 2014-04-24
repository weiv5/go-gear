package main

import (
    "gear"
    "web/app"
)

func main() {
    gear.StaticRoute()
    gear.AddRoute("/", &app.IndexApp{})
    gear.AddRoute("/test", &app.TestApp{})
    gear.Run()
}
