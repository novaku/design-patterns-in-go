package framework

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}