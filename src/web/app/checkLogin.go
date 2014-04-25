package app

import (
    "gear"
)

type CheckLogin struct {}

func (c *CheckLogin) Check(r *gear.Request) bool {
    if (r.Get("login") != "1") {
        return false
    }
    return true
}

func (c *CheckLogin) Failed(w *gear.Response) {
    w.Redirect("/login")
}
