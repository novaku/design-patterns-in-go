package main

import (
	"fmt"

	"design-patterns/structural-patterns/flyweight/framework"
)

func main() {
	game := framework.NewGame()

	//Add Terrorist
	game.AddTerrorist(framework.TerroristDressType)
	game.AddTerrorist(framework.TerroristDressType)
	game.AddTerrorist(framework.TerroristDressType)
	game.AddTerrorist(framework.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(framework.CounterTerroristDressType)
	game.AddCounterTerrorist(framework.CounterTerroristDressType)
	game.AddCounterTerrorist(framework.CounterTerroristDressType)

	dressFactoryInstance := framework.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}