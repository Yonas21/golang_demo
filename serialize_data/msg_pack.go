package serializedata

import (
	"github.com/vmihailenco/msgpack"
)

func MarshalData(p Person) ([]byte, error) {

	data, err := msgpack.Marshal(p)

	if err != nil {
		return nil, err

	} else {
		return data, nil
	}

}

func UnMarshalData(data []byte) (Person, error) {
	var person = Person{}
	err := msgpack.Unmarshal(data, &person)

	if err != nil {
		return Person{}, err
	} else {
		return person, nil
	}
}
