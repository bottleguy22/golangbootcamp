package main

import "fmt"

func main() {

	const (
		_   = iota
		Jan = 1
		Feb = 2
		Mar = 3
	)

	fmt.Println(Jan, Feb, Mar)
}
