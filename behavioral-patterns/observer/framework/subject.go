package framework

type Subject interface {
	Register(Observer Observer)
	Deregister(Observer Observer)
	NotifyAll()
}