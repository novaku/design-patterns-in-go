package main

import "design-patterns/behavioral-patterns/mediator/framework"

func main() {
	stationManager := framework.NewStationManger()

	passengerTrain := &framework.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &framework.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}