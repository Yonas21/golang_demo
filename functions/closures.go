/*
* This example demonstrate closure
 */
package functions

import (
	"fmt"
)

func getAnonymous() {
	input_NUM := 10
	squareNumber := func() int {
		input_NUM *= input_NUM
		return input_NUM
	}

	fmt.Println(squareNumber())
}
