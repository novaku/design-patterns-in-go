package framework

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}