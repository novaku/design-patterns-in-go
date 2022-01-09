package framework

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}