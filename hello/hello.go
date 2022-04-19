package main

import (
	. "fmt"

	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// slices of names
	names := []string{"Yonas", "mike", "Heni"}

	message, err := greetings.Hellos(names)

	// if an error was returned, print it to the console and
	// exit the program

	if err != nil {
		log.Fatal(err)
	}
	Println(message)

	Println(0)

}
