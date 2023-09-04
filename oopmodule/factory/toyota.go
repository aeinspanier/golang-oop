package factory

type Toyota struct {
    Car
}

func newToyota() ICar {
    return &Toyota{
        Car: Car{
            make: "Toyota",
			model: "RAV4",
        },
    }
}