package main

import (
	"fmt"

	"design-patterns/behavioral-patterns/visitor/framework"
)

func main() {
	square := &framework.Square{Side: 2}
	circle := &framework.Circle{Radius: 3}
	rectangle := &framework.Rectangle{L: 2, B: 3}

	areaCalculator := &framework.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &framework.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}