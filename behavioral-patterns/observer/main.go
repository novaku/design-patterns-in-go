package main

import "design-patterns/behavioral-patterns/observer/framework"

func main() {

	shirtItem := framework.NewItem("Nike Shirt")

	observerFirst := &framework.Customer{Id: "abc@gmail.com"}
	observerSecond := &framework.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}