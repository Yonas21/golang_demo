package main

import "fmt"

func main() {
	// get inputs from user
	fmt.Println("Enter Your Name: ")
	var myName string
	var randomPerson string

	fmt.Scanln(&myName)

	// get another name from the user
	fmt.Println("Get Another Person Name: ")

	fmt.Scanln(&randomPerson)

	if len(myName) <= len(randomPerson) {
		fmt.Println("That's Good Sign")
	} else if len(myName) == len(randomPerson) {
		fmt.Println("May be sharing the length is not bad")
	} else {
		fmt.Println("Option 3 is talking to a picture")
	}
}
