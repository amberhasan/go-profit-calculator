package main

import "fmt"

func main() {
    revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

    fmt.Println("%.1f", ebt)
	fmt.Println("%.1f", profit)
	fmt.Println("%.3f", ratio)
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Println(infoText)
	fmt.Scan(&userInput)
	return userInput
}