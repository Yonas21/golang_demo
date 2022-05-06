package observer

import . "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func NewItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) UpdateAvalability() {
	Printf("Item %s is now in Stock \n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *item) Register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) Deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}

	return observerList
}
