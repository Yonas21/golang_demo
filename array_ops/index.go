package main

import (
	. "fmt"
)

var numbers map[string]int

func main() {
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	// exists, 'ok' returns false. It returns true otherwise

	csharprating, ok := rating["C#"]

	if ok {
		Println("C# is in the map and its rating is", csharprating)
	} else {
		Println("We have no rating for C#")
	}

	delete(rating, "C")

	Println(rating)
}
