package main

import "fmt"

type duration int

// func (d *duration) pretty() string {
// 	return fmt.Sprintf("Duration: %d", *d)
// }

func (d duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func main() {
	duration(42).pretty()

	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}
