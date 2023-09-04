package main

import "fmt"
import "oopmodule/singleton"

func triggerSingleton() {
	for i := 0; i < 5; i++ {
        singleton.GetInstance()
    }
}

func main() {
	fmt.Println("Singleton: ")
	triggerSingleton();
}
