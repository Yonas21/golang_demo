package main

import (
	s "example/hello/serialize_data"
	"fmt"
	"log"
)

func main() {

	p := s.Person{Name: "Yonas Alem", Age: 25, Gender: "Male", Where: "Addis Ababa", Is_married: false}

	encoded, err := s.MarshalData(p)

	if err != nil {
		log.Fatalf("%v", err)
	}

	decoded, err := s.UnMarshalData(encoded)

	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(decoded)

}
