package main

import (
	"fmt"
	"go_in_action/chapter5/listing71/entities"
)

func main() {
	u := entities.User{
		Name: "Bill",
		//email: "bill@email.com",
	}

	fmt.Printf("User: %v\n", u)
}
