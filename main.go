package main

import (
	"example/hello/singleton"
	"fmt"
	"time"
)

func main() {

	// p := s.Person{Name: "Feven Fissiha", Age: 21, Gender: "Female", Place: "Addis Ababa", Is_married: true}

	// models.CreateUser(&p)
	func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*singleton.GetInstance())
	}()

	fmt.Println("welcome")

	for i := 0; i < 100; i++ {

		func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", singleton.GetInstance().Conn)
		}(i)
	}

}
