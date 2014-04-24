package app

import (
    "gear"
)

type CheckIp struct {}

func (c *CheckIp) Check(r *gear.Request) bool {
    if (r.Ip() == "127.0.0.1") {
        return false
    }
    return true
}

func (c *CheckIp) Failed(w *gear.Response) {
    w.Redirect("/")
}
