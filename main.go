package main

import (
	observer "example/hello/observer_pattern"
)

func main() {

	shirtItem := observer.NewItem("Jordan Air Shoes")

	observerOne := &observer.Customer{Id: "yonasalem056@gmail.com"}
	observerTwo := &observer.Customer{Id: "yonalem21@gmail.com"}

	shirtItem.Register(observerOne)
	shirtItem.Register(observerTwo)

	shirtItem.UpdateAvalability()

}
