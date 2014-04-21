package gear

import (
    "time"
)

type FuncM struct{}

func (f *FuncM) Date() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func (f *FuncM) Time() int64 {
    return time.Now().Unix();
}

var (
    Func = &FuncM{}
)
