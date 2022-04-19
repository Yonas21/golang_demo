package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	student_name := "Hello From Some Guest"
	fmt.Println(student_name)

	// check the length of string
	fmt.Println(reflect.TypeOf(student_name))

	// check the length of string
	fmt.Printf("the length of: %s, is, %d\n", student_name, len(student_name))

	// change the string to UpperCase
	fmt.Print(strings.ToUpper(student_name))
}
