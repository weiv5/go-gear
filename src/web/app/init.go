package app

import (
    "fmt"
    "gear"
)

type Verify struct {}

func (c *Verify) Check(r *gear.Request) bool {
    fmt.Println(r.Module)
    fmt.Println(r.Action)
    if (r.Ip() == "127.0.0.1") {
        return false
    }
    return true
}

func (c *Verify) Failed() {
    fmt.Println("failed")
}
