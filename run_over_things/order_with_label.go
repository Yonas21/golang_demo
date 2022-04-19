package main

import "fmt"

func main() {

	// define local variable
	var a int = 10

	/*loop over things and order to go*/
EXEC:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto EXEC
		}
		fmt.Printf("Value of a: %d \n", a)
		a++
	}
}
