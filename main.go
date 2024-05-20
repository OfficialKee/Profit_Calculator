package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {

	
	revenue,err1 := getRevenue()
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	expenses,err2 := getExpenses()
	if err2 != nil {
		fmt.Print(err2)
		return
	}
	taxRate,err3 := getTaxRate()
	if err3 != nil {
		fmt.Print(err3)
		return
	}
	
	

	earningsBeforeTax,profit,ratio := calculateEarnings(revenue,expenses,taxRate)
	writeToFile(earningsBeforeTax,profit,ratio)

	fmt.Printf("Earnings Before Tax: %v \nProfit: %v\nRatio: %v", earningsBeforeTax, profit, ratio)



}

func getRevenue()(float64,error){
	var revenue float64
	fmt.Println("What is your revenue? Enter below: ")
	fmt.Scan(&revenue)
	if revenue < 0 {
		return 0,errors.New("Revenue less than zero")
	}
	return revenue,nil
}
func getExpenses()(float64,error){
	var expenses float64
	fmt.Println("What are your expenses? Enter below: ")
	fmt.Scan(&expenses)
	if expenses < 0 {
		return 0,errors.New("Invalid expense amount")
	}
	return expenses,nil
}
func getTaxRate()(float64,error){
	var taxRate float64
	fmt.Println("What is your taxRate? Enter below: ")
	fmt.Scan(&taxRate)
	if taxRate < 0 {
		return 0,errors.New("Invalid tax rate")
	}
	return taxRate,nil
}
func calculateEarnings(revenue,expenses,taxRate float64)(float64,float64,float64){
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (earningsBeforeTax * (taxRate/100))
	ratio  := earningsBeforeTax / profit
	return earningsBeforeTax,profit,ratio
}

func writeToFile(earningsBeforeTax,profit,ratio float64){
	balanceText := fmt.Sprintf("EBT: %v \nProfit: %v \nRatio: %v",earningsBeforeTax,profit,ratio)
	os.WriteFile("balance.txt",[]byte(balanceText),0644)

}