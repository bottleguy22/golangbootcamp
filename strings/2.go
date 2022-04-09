/*string len, how to get
letters that you dont use in english, doesn't use only 1 byte*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "Inanç" //the result its 5 letters, but its showing 6 bytes because wasn't using utf8

	// fmt.Println(len(name))

	fmt.Println(utf8.RuneCountInString(name))
}

/*RuneCountInString = count the number of characters
should use the package unicode/utf8*/
