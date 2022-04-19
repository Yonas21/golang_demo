package main

import (
	"encoding/json"
	. "fmt"
	"log"
	"os"
)

func main() {

	map1 := map[string]string{
		"name": "Yonas Alem",
		"Job":  "Software Engineer",
	}

	file, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	value := "This is the Mighty Monday \n"
	data := []byte(value)

	_, err2 := file.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}

	message := "Greetings, Everybody"
	data2 := []byte(message)

	var _idx int64 = int64(len(data2) + 10)
	_, err3 := file.WriteAt(data2, _idx)

	if err3 != nil {
		log.Fatal(err3)
	}

	// convert map to json
	jsonStr, err4 := json.Marshal(map1)

	if err4 != nil {
		Printf("Error %s", err4)
	} else {
		Println(string(jsonStr))
	}

	Println("Done")

}
