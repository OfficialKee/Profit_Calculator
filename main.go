package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	// var earningsBeforeTax float64 = revenue - expenses
	// var profit float64 = revenue - ((revenue - expenses) * taxRate)
	// var ratio float64 = earningsBeforeTax / profit

	fmt.Println("What is your revenue? Enter below: ")
	fmt.Scan(&revenue)
	fmt.Println("What are your expenses? Enter below: ")
	fmt.Scan(&expenses)
	fmt.Println("What is your annual tax rate? Enter below: ")
	fmt.Scan(&taxRate)
	// fmt.Print("hi")
	// fmt.Sprintf("Earnings Before Tax: %v \nProfit: %v\nRatio: %v", earningsBeforeTax, profit, ratio)
	var earningsBeforeTax float64 = revenue - expenses
	var profit float64 = earningsBeforeTax - (earningsBeforeTax * (taxRate/100))
	var ratio float64 = earningsBeforeTax / profit

	fmt.Printf("Earnings Before Tax: %v \nProfit: %v\nRatio: %v", earningsBeforeTax, profit, ratio)

}
