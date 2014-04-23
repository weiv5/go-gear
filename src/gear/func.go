package gear

import (
    "time"
)

func Date() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func Time() int64 {
    return time.Now().Unix();
}
