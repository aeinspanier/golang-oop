package main

import "fmt"
import "oopmodule/singleton"
import "oopmodule/factory"
import . "oopmodule/adapter"
import . "oopmodule/iterator"

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

func runAdapter() {
	client := &Client{}
    mac := &Mac{}

    client.InsertLightningConnectorIntoComputer(mac)

    windowsMachine := &Windows{}
    windowsMachineAdapter := &WindowsAdapter{
        WindowMachine: windowsMachine,
    }

    client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

func runIterator() {
	user1 := &User{
        Name: "a",
        Age:  30,
    }
    user2 := &User{
        Name: "b",
        Age:  20,
    }

    userCollection := &UserCollection{
        Users: []*User{user1, user2},
    }

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)

	}
}

func main() {
	fmt.Println("Singleton: ")
	triggerSingleton()

	fmt.Println("Factory: ")
	runFactory()

	fmt.Println("Adapter: ")
	runAdapter()

	fmt.Println("Iterator: ")
	runIterator()
}
