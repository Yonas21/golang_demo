package runoverthings

import "fmt"

func getRange() {

	b := 15
	var a int

	numbers := [6]int{1, 2, 3, 4, 10}

	// let run over some numbers
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("Value of a: %d", a)
	}

	for i, x := range numbers {
		fmt.Printf("Value of X = %d at %d \n", x, i)
	}

}
