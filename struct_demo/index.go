package main

import (
	"fmt"
	"time"
)

//defining struct type
type Transaction struct {
	billNo       int32
	created_date time.Time
	payerName    string
}

//embedd one kind of struct to another

type Report struct {
	Transaction
	paid_date time.Time
}

func (t Transaction) getTransactionDetails() Transaction {
	return t
}

func (r Report) getReportDetails() Report {
	return r
}

func main() {
	print := fmt.Println
	payer := Transaction{billNo: 45334, created_date: time.Now(), payerName: "Yonas Alem"}

	report1 := Report{Transaction: Transaction{billNo: 45334, created_date: time.Now(), payerName: "Yonas Alem"}, paid_date: time.Now()}

	print(payer.getTransactionDetails())

	print(payer.created_date.Format(time.ANSIC))

	print(report1.getReportDetails().paid_date.Format(time.RFC1123))
}
