package framework

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstanceOne *single

func GetInstanceOne() *single {
	if singleInstanceOne == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstanceOne = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstanceOne
}