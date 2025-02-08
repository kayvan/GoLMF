package main

import (
	"fmt"
)

func tipCalculator() {
	var bill, tip float64

	fmt.Print("Enter the bill amount: €")
	fmt.Scanln(&bill)

	fmt.Print("Enter the tip percentage: ")
	fmt.Scanln(&tip)

	tipAmount := bill * (tip / 100)
	total := bill + tipAmount

	fmt.Printf("Tip amount: %.2f€\nTotal amount: %.2f€\n", tipAmount, total)
}
