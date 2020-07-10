package main

import (
	"fmt"
	"go_in_action/chapter5/listing64/counters"
)

func main() {
	//counter := counters.alertCounter(10)
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
