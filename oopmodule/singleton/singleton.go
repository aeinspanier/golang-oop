package singleton

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleTon *single

func createSingleton() {
	fmt.Println("Create singleton")
	singleTon = &single{}
}

func GetInstance() *single{
	if singleTon != nil {
		fmt.Println("singleton already exists")
	} else {
		once.Do(createSingleton)
	}
	return singleTon
}
