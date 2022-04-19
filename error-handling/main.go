package main

import (
	. "fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")

	if user == "" {
		panic("no value for $USER")
	} else {
		Println("$USER")
	}
}
