package gear

import (
    "net/http"
    "strconv"
    "strings"
    "fmt"
    "html/template"
    "encoding/json"
)

type AppInterface interface {
    IndexAction()
}

type App struct {
    Module string
    Action string
    R *http.Request
    W http.ResponseWriter
    Data map[string] interface{}
}

func (app *App) Init(w http.ResponseWriter, r *http.Request, module string, action string) {
    app.W = w
    app.R = r
    app.Data = make(map[string] interface{})
    app.Module = module
    app.Action = action
}

func (app *App) Before() int {
    return 123
}

func (app *App) Display(name string, tpl ...string) error {
    t,_ := template.ParseFiles(tpl...)
    t.ExecuteTemplate(app.W, name, app.Data)
    t.Execute(app.W, nil)
    return nil
}

func (app *App) Json(data interface{}) error {
    content,_ := json.Marshal(data)
    app.W.Header().Set("Content-Type", "application/json;charset=UTF-8")
    app.W.Header().Set("Content-Length", strconv.Itoa(len(content)))
    app.W.Write(content)
    return nil
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
    return app.R.FormValue(name)
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
    return app.R.PostFormValue(name)
}

func (app *App) Ip() string {
    ips := app.R.Header.Get("X-Forwarded-For")
    if ips != "" {
        ip := strings.Split(ips, ",")
        if len(ip) > 0 && ip[0] != "" {
             return ip[0]
        }
    }
    ips2 := strings.Split(app.R.RemoteAddr, ":")
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
