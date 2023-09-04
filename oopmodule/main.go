package main

import "fmt"
import "oopmodule/singleton"
import "oopmodule/factory"

func triggerSingleton() {
	for i := 0; i < 5; i++ {
        singleton.GetInstance()
    }
}

func runFactory() {
	toyota, _ := factory.GetCar("Toyota")
	honda, _ := factory.GetCar("Honda")
	fmt.Println("Checking Honda: Make=" + honda.GetMake() + " Model=" + honda.GetModel())
	fmt.Println("Checking Toyota: Make=" + toyota.GetMake() + " Model=" + toyota.GetModel())
}

func main() {
	fmt.Println("Singleton: ")
	triggerSingleton()

	fmt.Println("Factory: ")
	runFactory()
}
