package main

import (
	"fmt"
)

func main() {
	/*local variable definition*/
	var grade string
	var marks int

	// get both inputs from Keyboard
	fmt.Print("Please Enter Your Score: ")

	fmt.Scanln(&marks)

	average := int(marks / 10)
	switch average {
	case 9:
		grade = "A"
	case 8:
		grade = "B"
	case 5, 6, 7:
		grade = "C"
	case 4:
		grade = "D"
	default:
		grade = "F"
	}

	// build conditions over marks and add some remarks
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}

	fmt.Printf("Your Grade is %s \n", grade)

	fmt.Println("Check the Average", average)

}
