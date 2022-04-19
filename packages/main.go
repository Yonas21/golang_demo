package main

import (
	"example/hello/packages/strutil"
	"fmt"
	"math"
	// "github.com/Yonas21/getting-started-with-go/packages/strutil"
)

func main() {
	fmt.Println(math.Round(4.74))
	fmt.Println(math.Ceil(5.7))

	fmt.Printf("The Square root of 64 is %v \n", math.Sqrt(64))

	fmt.Println(strutil.Reverse("olleH"))
}
