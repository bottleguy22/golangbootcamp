package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	l := len(msg)

	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)

	fmt.Println(s)
}

/* using os.Args, it will wait for the args after the go run 3.go
to upper will use to put in capital
strings.Repeat will repeat the string and the signals that you use */
