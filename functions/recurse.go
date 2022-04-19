package functions

import "fmt"

func printFactorial() {
	var input_num int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&input_num)

	fmt.Printf("The Factorial of ,%d, is, %d", input_num, factorial(input_num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
