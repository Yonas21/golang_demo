package functions

import "fmt"

type Employee struct {
	fname string
	lname string
}

type Numbers struct {
	number1 int32
	number2 int32
}

func (emp Employee) fullName() {
	fmt.Println(emp.fname + ", " + emp.lname)
}

func (num Numbers) getSum() int32 {
	return num.number1 + num.number2
}

func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0

	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value

	}

	return finalAddValue, finalSubValue
}

func printStruct() {
	employee1 := Employee{"Yonas", "Alem"}

	numbers := Numbers{34, 67}

	employee1.fullName()

	fmt.Printf("The sum of %d, and %d,is, %d \n", numbers.number1, numbers.number2, numbers.getSum())

	//sum of all values
	fmt.Println(addAll(1, 2, 34, 5, 1, 54, 22))
}
