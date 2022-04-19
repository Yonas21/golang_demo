package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30

	random_numbers := []int{2, 45}

	for i := 0; i < 1000000; i++ {

		rand_number := rand.Intn(max-min) + min

		random_numbers = append(random_numbers, rand_number)
	}
	fmt.Println(random_numbers)

}
