package main

import (
	"example/hello/models"
	"fmt"
	"log"
)

func main() {

	// p := s.Person{Name: "Feven Fissiha", Age: 21, Gender: "Female", Place: "Addis Ababa", Is_married: true}

	// models.CreateUser(&p)

	users, err := models.GetAllUsers()

	if err != nil {
		log.Fatalf("unable to get any user: %v", err)
	}

	for _, v := range users {
		fmt.Printf("%+v\n", v)

	}

}
