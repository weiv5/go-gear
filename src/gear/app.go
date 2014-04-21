package gear

import (
    "net/http"
    "strconv"
    "strings"
    "fmt"
)

type AppInterface interface {
    IndexAction()
}

type App struct {
    r *http.Request
    w http.ResponseWriter
}

func (app *App) Init(w http.ResponseWriter, r *http.Request) {
    app.w = w
    app.r = r
}

func (app *App) GetInt(name string) (int, error) {
    return strconv.Atoi(app.Get(name))
}

func (app *App) GetInt64(name string) (int64, error) {
    return strconv.ParseInt(app.Get(name), 10, 64)
}

func (app *App) GetFloat(name string) (float64, error) {
    return strconv.ParseFloat(app.Get(name), 64)
}

func (app *App) Get(name string) string {
    return app.r.FormValue(name)
}

func (app *App) PostInt(name string) (int, error) {
    return strconv.Atoi(app.Post(name))
}

func (app *App) PostInt64(name string) (int64, error) {
    return strconv.ParseInt(app.Post(name), 10, 64)
}

func (app *App) PostFloat(name string) (float64, error) {
    return strconv.ParseFloat(app.Post(name), 64)
}

func (app *App) Post(name string) string {
    return app.r.PostFormValue(name)
}

func (app *App) Ip() string {
    ips := app.r.Header.Get("X-Forwarded-For")
    if ips != "" {
        ip := strings.Split(ips, ",")
        if len(ip) > 0 && ip[0] != "" {
             return ip[0]
        }
    }
    ips2 := strings.Split(app.r.RemoteAddr, ":")
    if len(ips2) > 0 {
        if ips2[0] != "[" {
            return ips2[0]
        }
    }
    return "127.0.0.1"
}

func (app *App) Ip2Long(ip ...string) int64 {
    var ipStr string
    if len(ip) == 0 {
        ipStr = app.Ip()
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
