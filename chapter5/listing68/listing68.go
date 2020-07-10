package main

import (
	"fmt"	
	"go_in_action/chapter5/listing68/counters"
)


func main() {
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n",counter)
}