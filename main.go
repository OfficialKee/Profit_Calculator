package main

import (
	"fmt"
)

func main() {

	
	revenue := getRevenue()
	expenses := getExpenses()
	taxRate := getTaxRate()

	earningsBeforeTax,profit,ratio := calculateEarnings(revenue,expenses,taxRate)

	fmt.Printf("Earnings Before Tax: %v \nProfit: %v\nRatio: %v", earningsBeforeTax, profit, ratio)

}

func getRevenue()float64{
	var revenue float64
	fmt.Println("What is your revenue? Enter below: ")
	fmt.Scan(&revenue)
	return revenue
}
func getExpenses()float64{
	var expenses float64
	fmt.Println("What are your expenses? Enter below: ")
	fmt.Scan(&expenses)
	return expenses
}
func getTaxRate()float64{
	var taxRate float64
	fmt.Println("What is your taxRate? Enter below: ")
	fmt.Scan(&taxRate)
	return taxRate
}
func calculateEarnings(revenue,expenses,taxRate float64)(float64,float64,float64){
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate/100))
	ratio  := earningsBeforeTax / profit
	return earningsBeforeTax,profit,ratio
}