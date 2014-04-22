package main

import (
    "gear"
    "web/app"
)


func main() {
    gear.AddRoute("/", &app.IndexApp{})
    gear.AddRoute("/test", &app.TestApp{})
    gear.Run()
}
