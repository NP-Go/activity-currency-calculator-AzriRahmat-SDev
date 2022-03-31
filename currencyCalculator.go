package main

import (
	"fmt"
	"math"
)

var totalAmount float64
var twoDollarNotes int
var changeAmount float64

func CurrencyCalculator(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin float64) (float64, int, float64) {
	//Insert your code here
	totalAmount = oneDollarCoin + (fiftyCentCoin * 0.50) + (twentyCentCoin * 0.20) + (tenCentCoin * 0.10) + (fiveCentCoin * 0.05)
	twoDollarNotes = int(totalAmount) / 2
	changeAmount = math.Mod(totalAmount, 1)
	//Do not remove this
	fmt.Printf("Total: %0.2f \n", totalAmount)
	fmt.Println("Two Dollar Notes:", twoDollarNotes)
	fmt.Printf("Change: %0.2f \n", changeAmount)
	return totalAmount, twoDollarNotes, changeAmount
}

func main() {
	var oneDollarCoin float64
	var fiftyCentCoin float64
	var twentyCentCoin float64
	var tenCentCoin float64
	var fiveCentCoin float64

	fmt.Println("Enter the number of 1-dollar coins: ")
	fmt.Scanln(&oneDollarCoin)
	fmt.Println("Enter the number of 50-cent coins: ")
	fmt.Scanln(&fiftyCentCoin)
	fmt.Println("Enter the number of 20-cent coins: ")
	fmt.Scanln(&twentyCentCoin)
	fmt.Println("Enter the number of 10-cent coins: ")
	fmt.Scanln(&tenCentCoin)
	fmt.Println("Enter the number of 5-cent coins: ")
	fmt.Scanln(&fiveCentCoin)
	CurrencyCalculator(oneDollarCoin, fiftyCentCoin, twentyCentCoin, tenCentCoin, fiveCentCoin)
}
