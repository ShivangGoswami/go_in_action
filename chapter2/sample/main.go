package main


import (
	"log"
	"os"

	_ "go_in_action/chapter2/sample/matchers"
	"go_in_action/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main(){
	search.Run("president")
}