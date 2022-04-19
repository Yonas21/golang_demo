package main

import (
	s "example/hello/serialize_data"
	"fmt"
	"log"
)

func main() {

	p := s.Person{"satori", 16, "F", "Oriental spirit hall", false}

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
