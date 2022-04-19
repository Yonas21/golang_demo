/*
* This example demonstrate closure
 */
package main

import (
	"fmt"
)

func main() {
	input_NUM := 10
	squareNumber := func() int {
		input_NUM *= input_NUM
		return input_NUM
	}

	fmt.Println(squareNumber())
}
