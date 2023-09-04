package factory

import "fmt"

func GetCar(carType string) (ICar, error) {
    if carType == "Honda" {
        return newHonda(), nil
    }
    if carType == "Toyota" {
        return newToyota(), nil
    }
    return nil, fmt.Errorf("Invalid car type")
}