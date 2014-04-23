package gear

type CheckInterface interface {
    Check() bool
    Failed()
}
