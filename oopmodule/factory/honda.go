package factory

type Honda struct {
    Car
}

func newHonda() ICar {
    return &Honda{
        Car: Car{
            make: "Honda",
			model: "CR-V",
        },
    }
}