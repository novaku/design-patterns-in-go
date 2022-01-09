package main

import (
	"fmt"

	"design-patterns/creational-patterns/factory-method/framework"
)

func main() {
	ak47, _ := framework.GetGun("ak47")
	musket, _ := framework.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g framework.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}