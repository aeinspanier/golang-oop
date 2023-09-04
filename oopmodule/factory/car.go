package factory

type Car struct {
    make  string
    model string
}

func (c *Car) setMake(make string) {
    c.make = make
}

func (c *Car) GetMake() string {
    return c.make
}

func (c *Car) setModel(model string) {
    c.model = model
}

func (c *Car) GetModel() string {
    return c.model
}