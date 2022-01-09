package framework

type Collection interface {
	CreateIterator() Iterator
}