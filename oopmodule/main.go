package main

import "fmt"
import "os"
import "oopmodule/singleton"
import "oopmodule/factory"
import "github.com/akamensky/argparse"
import . "oopmodule/adapter"
import . "oopmodule/iterator"
import . "oopmodule/builder"
import . "oopmodule/observer"

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

func runBuilder() {
	normalBuilder := GetBuilder("normal")
	iglooBuilder := GetBuilder("igloo")

    director := NewDirector(normalBuilder)
    normalHouse := director.BuildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}

func runObserver() {
	shirtItem := NewItem("Nike Shirt")

    observerFirst := &Customer{Id: "abc@gmail.com"}
    observerSecond := &Customer{Id: "xyz@gmail.com"}

    shirtItem.Register(observerFirst)
    shirtItem.Register(observerSecond)

    shirtItem.UpdateAvailability()
}

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	var singletonFlag *bool = parser.Flag("s", "singleton", &argparse.Options{Required: false, Help: "Flag to execute singleton"})
	var factoryFlag *bool = parser.Flag("f", "factory", &argparse.Options{Required: false, Help: "Flag to execute factory"})
	var adapterFlag *bool = parser.Flag("a", "adapter", &argparse.Options{Required: false, Help: "Flag to execute adapter"})
	var iteratorFlag *bool = parser.Flag("i", "iterator", &argparse.Options{Required: false, Help: "Flag to execute iterator"})
	var builderFlag *bool = parser.Flag("b", "builder", &argparse.Options{Required: false, Help: "Flag to execute builder"})
	var observerFlag *bool = parser.Flag("o", "observer", &argparse.Options{Required: false, Help: "Flag to execute observer"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}


	if *singletonFlag {
		fmt.Println("Singleton: ")
		triggerSingleton()
	}

	if *factoryFlag {
		fmt.Println("Factory: ")
		runFactory()
	}

	if *adapterFlag {
		fmt.Println("Adapter: ")
		runAdapter()
	}

	if *iteratorFlag {
		fmt.Println("Iterator: ")
		runIterator()
	}

	if *builderFlag {
		fmt.Println("Builder: ")
		runBuilder()
	}

	if *observerFlag {
		fmt.Println("Observer: ")
		runObserver()
	}
	
}
