package framework

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (g *PassengerTrain) Arrive() {
	if !g.Mediator.CanArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.NotifyAboutDeparture()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.Arrive()
}