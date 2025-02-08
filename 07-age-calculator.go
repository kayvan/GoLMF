package main

import (
	"fmt"
	"time"
)

func ageCalculator() {
	var birthYear int
	fmt.Print("Enter your birth year: ")
	fmt.Scanln(&birthYear)
	currentYear := time.Now().Year()
	fmt.Println("Current year:", currentYear)
	age := currentYear - birthYear
	fmt.Printf("You are %d years old!\n", age)
}