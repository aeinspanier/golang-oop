package factory

type ICar interface {
    setMake(make string)
    setModel(model string)
    GetMake() string
    GetModel() string
}