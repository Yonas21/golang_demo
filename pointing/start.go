package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 45
	name := "Yonas"
	changeValue(&x)
	fmt.Println(x)

	fmt.Println(reflect.TypeOf(name))
}

func changeValue(value *int) {
	*value = 0
}
