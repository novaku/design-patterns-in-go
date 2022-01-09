package framework

type Shape interface {
	GetType() string
	Accept(Visitor)
}