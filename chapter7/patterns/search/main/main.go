package main

import (
	"go_in_action/chapter7/patterns/search"

	"log"
)

func main() {
	results := search.Submit(
		"golang",
		search.OnlyFirst,
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}

	results = search.Submit(
		"golang",
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}
}
