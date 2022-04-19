package main

import "fmt"

func main() {
	// shorthand declaration of variable
	studentName := "Yonas"
	//declare and assign variables
	// var studentName = "Yonas"

	var age int64 = 25

	const isAdult = true

	height := 1.74

	// isAdult = false

	// another shorthand of declaring two in one
	name, email := "Yonas Alem", "yonalem21@gmail.com"

	// print out the name
	fmt.Println(studentName, age, isAdult)

	// print out
	fmt.Println(name, email)

	// use string formatting
	fmt.Printf("%T \n", height)
}
