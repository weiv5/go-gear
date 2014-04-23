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

func (r *Response) Display(name string, tpl ...string) error {
    t,_ := template.ParseFiles(tpl...)
    t.ExecuteTemplate(r.W, name, r.Data)
    t.Execute(r.W, nil)
    return nil
}

func (r *Response) Json(data interface{}) error {
    content,_ := json.Marshal(data)
    r.W.Header().Set("Content-Type", "application/json;charset=UTF-8")
    r.W.Header().Set("Content-Length", strconv.Itoa(len(content)))
    r.W.Write(content)
    return nil
}
