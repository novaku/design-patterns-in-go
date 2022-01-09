package framework

type Iterator interface {
	HasNext() bool
	GetNext() *User
}