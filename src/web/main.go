package main

import (
    "gear"
    "web/app"
)

func main() {
    gear.AddRoute("/test", &app.TestApp{})
    gear.Run()
}
