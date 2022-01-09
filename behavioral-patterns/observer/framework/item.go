package framework

import "fmt"

type Item struct {
	ObserverList []Observer
	Name         string
	InStock      bool
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}
func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}
func (i *Item) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

func (i *Item) Deregister(o Observer) {
	i.ObserverList = RemoveFromSlice(i.ObserverList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func RemoveFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}