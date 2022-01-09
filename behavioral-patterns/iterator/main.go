package main

import (
	"fmt"

	"design-patterns/structural-patterns/iterator/framework"
)

func main() {

	user1 := &framework.User{
		Name: "a",
		Age:  30,
	}
	user2 := &framework.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &framework.UserCollection{
		Users: []*framework.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}