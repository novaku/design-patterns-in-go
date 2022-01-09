package main

import (
	"fmt"

	"design-patterns/creational-patterns/singleton/framework"
)

func main() {
	// try one of this
	conceptualExample()

	// or
	// anotherExample()
}

func conceptualExample() {
	for i := 0; i < 30; i++ {
		go framework.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

func anotherExample() {
	for i := 0; i < 30; i++ {
		go framework.GetInstanceOne()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}