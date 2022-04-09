package main

import "fmt"

func main() {

	const (
		Nov = 11 - iota
		Oct = 10
		Sep = 9
	)

	fmt.Println(Sep, Oct, Nov)
}
