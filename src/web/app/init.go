package app

import (
    "fmt"
)

type Verify struct {}

func (c *Verify) Check() bool {
    return false
}

func (c *Verify) Failed() {
    fmt.Println("failed")
}
