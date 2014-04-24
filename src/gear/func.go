package gear

import (
    "time"
    "strings"
)

func Date(f ...string) string {
    var format string
    if len(f)==1 {
        format = strings.ToLower(f[0])
        format = strings.Replace(format, "y", "2006", 1)
        format = strings.Replace(format, "m", "01", 1)
        format = strings.Replace(format, "d", "02", 1)
        format = strings.Replace(format, "h", "15", 1)
        format = strings.Replace(format, "i", "04", 1)
        format = strings.Replace(format, "s", "05", 1)
    } else {
        format = "2006-01-02 15:04:05"
    }
    return time.Now().Format(format)
}

func Time() int64 {
    return time.Now().Unix();
}
