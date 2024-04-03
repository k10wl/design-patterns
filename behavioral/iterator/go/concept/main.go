package main

import (
	"fmt"
)

func main() {
	fmt.Println(">>> Iterator demo")
	users := UserCollection{
		users: []User{
			{id: 1, name: "Rick"},
			{id: 2, name: "Marty"},
			{id: 3, name: "Gloria"},
			{id: 4, name: "Helen"},
		},
	}
	logAllUsers(users.createUserIterator())
}

func logAllUsers(iterator Iterator) {
	// I don't care what comes in and how the iteration happens
	fmt.Println("Logging user data:")
	for iterator.hasMore() {
		fmt.Printf("> %+v\n", iterator.getNext())
	}
}
