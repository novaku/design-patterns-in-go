package main

import (
	"fmt"

	"design-patterns/structural-patterns/decorator/framework"
)

func main() {

	pizza := &framework.VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &framework.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &framework.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())

	peppyPaneerPizza := &framework.PeppyPaneer{}

	//Add cheese topping
	peppyPaneerPizzaWithCheese := &framework.CheeseTopping{
		Pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.GetPrice())
}