package gear

import (
    "reflect"
    "strings"
    "net/http"
    "os"
)

var (
    StaticMaps  = make(map[string]string)
    RouterMaps  = make(map[string]map[string]reflect.Type)
)

func StaticRoute() {
    dir := strings.TrimSuffix(Ini.String("static::dir"), "/")+"/"
    js := Ini.String("static::js")
    if js!="" {
        StaticMaps["js"] = dir+strings.Trim(js, "/")
    }
    css := Ini.String("static::css")
    if css!="" {
        StaticMaps["css"] = dir+strings.Trim(css, "/")
    }
    image := Ini.String("static::image")
    if image!="" {
        StaticMaps["image"] = dir+strings.Trim(image, "/")
    }
    favicon := Ini.String("static::favicon.ico")
    if favicon!="" {
        StaticMaps["favicon.ico"] = dir+strings.Trim(favicon, "/")
    }
}

func AddRoutes(path string, apps ...AppInterface) {
    if len(apps) == 0 {
        return
    }
    for _,v := range apps {
        AddRoute(path, v)
    }
}

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
            action := strings.TrimSuffix(p.Method(i).Name, "Action")
            if _,ok := RouterMaps[m][action]; !ok && action!="" {
                RouterMaps[m][action] = t
            }
        }
    }
}

type Serve struct {}
func (serve *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    request := &Request{R:r}
    response:= &Response{W:w}
    Log.Access(request)

    path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    var m, action string
    l := len(path)
    //static file
    if l>0 {
        if static, ok := StaticMaps[path[0]]; ok {
            file := static + strings.TrimPrefix(strings.Trim(r.URL.Path, "/"), path[0])
            finfo, err := os.Stat(file)
            if err != nil || finfo.IsDir() {
                http.NotFound(w, r)
            } else {
                http.ServeFile(w, r, file)
            }
            return
        }
    }
    //auto route
    if l==0 {
        m, action = "", "Index"
    } else if l==1 {
        m, action = path[0], "Index"
    } else if l==2 {
        m, action = path[0], strings.Title(path[1])
    }
    if appType, ok := RouterMaps[m][action]; ok {
        request.Module = m
        request.Action = action
        app := reflect.New(appType)

        checkType := reflect.TypeOf((*CheckInterface)(nil)).Elem()
        if app.Type().Implements(checkType) {
            check := app.MethodByName("Check")
            checkRes := check.Call([]reflect.Value{reflect.ValueOf(request)})
            if checkRes[0].Bool() == false {
                failed := app.MethodByName("Failed")
                failed.Call([]reflect.Value{reflect.ValueOf(response)})
                return
            }
        }

        init := app.MethodByName("Init")
        init.Call([]reflect.Value{reflect.ValueOf(response), reflect.ValueOf(request)})

        method := app.MethodByName(action+"Action")
        method.Call(nil)
        return
    }
    http.NotFound(w, r)
}
