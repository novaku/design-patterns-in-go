package framework

type Observer interface {
	Update(string)
	GetID() string
}