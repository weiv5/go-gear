package gear

import (
    "net/http"
    "html/template"
    "encoding/json"
    "strconv"
)

type Response struct {
    Data map[string] interface{}
    W http.ResponseWriter
}

func (w *Response) SetHeader(key string, val string) {
    w.W.Header().Set(key, val);
}

func (w *Response) Display(name string, tpl ...string) error {
    t,_ := template.ParseFiles(tpl...)
    t.ExecuteTemplate(w.W, name, w.Data)
    t.Execute(w.W, nil)
    return nil
}

func (w *Response) Json(data interface{}) error {
    content,_ := json.Marshal(data)
    w.SetHeader("Content-Type", "application/json;charset=UTF-8")
    w.SetHeader("Content-Length", strconv.Itoa(len(content)))
    w.W.Write(content)
    return nil
}

func (w *Response) Redirect(url string) {
    w.SetHeader("Location", url)
    w.W.WriteHeader(307)
}
