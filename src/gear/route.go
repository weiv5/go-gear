package gear

import (
    "reflect"
    "strings"
    "net/http"
    "fmt"
)

var (
    RouterMaps  = make(map[string]map[string]reflect.Type)
)

func AddRoute(path string, app AppInterface) {
    v := reflect.ValueOf(app)
    p := reflect.TypeOf(app)
    t := reflect.Indirect(v).Type()
    m := strings.Trim(path, "/")

    if _, ok := RouterMaps[m]; !ok {
        RouterMaps[m] = make(map[string]reflect.Type)
    }
    for i := 0; i < p.NumMethod(); i++ {
        if strings.HasSuffix(p.Method(i).Name, "Action") {
            action := strings.ToLower(strings.TrimSuffix(p.Method(i).Name, "Action"))
            if _,ok := RouterMaps[m][action]; !ok && action!="" {
                RouterMaps[m][action] = t
            }
        }
    }
}

type Serve struct {}
func (serve *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(strings.ToLower(strings.Trim(r.URL.Path, "/")), "/")
    var m, action string
    l := len(path)
    if l==0 {
        m, action = "", "index"
    } else if l==1 {
        m, action = path[0], "index"
    } else if l==2 {
        m, action = path[0], path[1]
    }
    if appType, ok := RouterMaps[m][action]; ok {
        app := reflect.New(appType)

        init := app.MethodByName("Init")
        init.Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r), reflect.ValueOf(m), reflect.ValueOf(action)})

        checkType := reflect.TypeOf((*CheckInterface)(nil)).Elem()
        if app.Type().Implements(checkType) {
            request := &Request{r}
            check := app.MethodByName("Check")
            checkRes := check.Call([]reflect.Value{reflect.ValueOf(request)})
            if checkRes[0].Bool() == false {
                failed := app.MethodByName("Failed")
                failed.Call(nil)
                return
            }
        }

        method := app.MethodByName(strings.Title(action)+"Action")
        method.Call(nil)
        return
    }
    serve.NotFound(w)
}

func (serve *Serve) NotFound(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(404)
    fmt.Fprintln(w, "404 page not found")
    return
}
