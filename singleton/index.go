package singleton

import (
	"sync"
)

var once sync.Once

type single struct {
	Conn string
}

var singleInstance *single

func GetInstance() *single {
	// if singleInstance == nil {

	// 	if singleInstance == nil {
	// 		fmt.Println("Creating single new Instance")
	// 		singleInstance = &single{}
	// 	} else {
	// 		fmt.Println("Instance Already Created")
	// 	}
	// } else {
	// 	fmt.Println("Singe Instance Already Created")
	// }

	once.Do(func() {
		singleInstance = &single{Conn: "Database Connected Once"}
	})
	return singleInstance
}
