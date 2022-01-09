package framework

import "fmt"

type MiddleCoordinates struct {
	X int
	Y int
}

func (a *MiddleCoordinates) VisitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) VisitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}