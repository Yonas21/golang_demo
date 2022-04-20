package serializedata

import (
	"bytes"
	"encoding/gob"
)

type Person struct {
	Name       string
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Place      string `json:"place"`
	Is_married bool   `json:"is_married"`
	Id         int    `json:"id"`
}

func EncodeData(data Person) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	if err := enc.Encode(data); err != nil {
		return nil, err
	} else {
		return buf, nil
	}
}

func DecodeData(buf1 *bytes.Buffer) (Person, error) {
	var another_girl = Person{}
	dec := gob.NewDecoder(bytes.NewBuffer(buf1.Bytes()))

	if err := dec.Decode(&another_girl); err != nil {
		return Person{}, err
	} else {
		return another_girl, nil
	}

}
