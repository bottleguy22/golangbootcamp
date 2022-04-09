/* blank identifies
timezones est, mst, pst
est = east standard time -5
mst = mounter standard time -7
pst = california standard time -8
*/
package main

import "fmt"

func main() {

	const (
		EST = -(5 + 0)
		_
		MST = -(5 + 2)
		PST = -(5 + 3)

		//est = -5
		//mst = -7
		//pst = -8
	)
	fmt.Println(EST, MST, PST)
}
