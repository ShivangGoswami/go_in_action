package main

import (
	"fmt"
	"go_in_action/chapter5/listing74/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}