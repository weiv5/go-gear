package gear

type CheckInterface interface {
    Check(r *Request) bool
    Failed()
}
