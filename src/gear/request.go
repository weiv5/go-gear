package gear

import (
    "net/http"
    "strings"
    "strconv"
    "fmt"
)

type Request struct {
    R *http.Request
    Module string
    Action string
}

func (r *Request) GetInt(name string) (int, error) {
    return strconv.Atoi(r.Get(name))
}

func (r *Request) GetInt64(name string) (int64, error) {
    return strconv.ParseInt(r.Get(name), 10, 64)
}

func (r *Request) GetFloat(name string) (float64, error) {
    return strconv.ParseFloat(r.Get(name), 64)
}

func (r *Request) Get(name string) string {
    return r.R.FormValue(name)
}

func (r *Request) PostInt(name string) (int, error) {
    return strconv.Atoi(r.Post(name))
}

func (r *Request) PostInt64(name string) (int64, error) {
    return strconv.ParseInt(r.Post(name), 10, 64)
}

func (r *Request) PostFloat(name string) (float64, error) {
    return strconv.ParseFloat(r.Post(name), 64)
}

func (r *Request) Post(name string) string {
    return r.R.PostFormValue(name)
}

func (r *Request) Ip() string {
    ips := r.R.Header.Get("X-Forwarded-For")
    if ips != "" {
        ip := strings.Split(ips, ",")
        if len(ip) > 0 && ip[0] != "" {
             return ip[0]
        }
    }
    ips2 := strings.Split(r.R.RemoteAddr, ":")
    if len(ips2) > 0 {
        if ips2[0] != "[" {
            return ips2[0]
        }
    }
    return "127.0.0.1"
}

func (r *Request) Ip2Long(ip ...string) int64 {
    var ipStr string
    if len(ip) == 0 {
        ipStr = r.Ip()
    } else {
        ipStr = ip[0]
    }
    ip_pieces := strings.Split(ipStr, ".")
    ip_1,_ := strconv.ParseInt(ip_pieces[0], 10, 32)
    ip_2,_ := strconv.ParseInt(ip_pieces[1], 10, 32)
    ip_3,_ := strconv.ParseInt(ip_pieces[2], 10, 32)
    ip_4,_ := strconv.ParseInt(ip_pieces[3], 10, 32)

    ip_bin := fmt.Sprintf("%08b%08b%08b%08b", ip_1, ip_2, ip_3, ip_4)
    ip_int,_ := strconv.ParseInt(ip_bin, 2, 64)
    return ip_int;
}
