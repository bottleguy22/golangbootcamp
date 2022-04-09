package main

import "fmt"

func main() {

	const (
		_      = iota
		Winter = 12 * (iota + 7)
		Spring = 3
		Summer = 6
		Fall   = 9
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}

//we stablish that iota its using 1, but can be any value inside the brackets
