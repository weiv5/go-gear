package gear

import (
    "reflect"
    "strings"
    "net/http"
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
    if l==1 {
        m, action = path[0], "index"
    } else if l==2 {
        m, action = path[0], path[1]
    }
    if appType, ok := RouterMaps[m][action]; ok {
        app := reflect.New(appType)

        wr := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
        init := app.MethodByName("Init")
        init.Call(wr)

        method := app.MethodByName(strings.Title(action)+"Action")
        method.Call(nil)
        return
    }
}
