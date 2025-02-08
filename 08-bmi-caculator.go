package main

import "fmt"


func BMI_Calculator() {
	var weight, height float64

	fmt.Print("Enter your weight in kilograms: ")
	fmt.Scanln(&weight)

	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&height)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI is: %.2f\n", bmi)
	if bmi < 18.5 {
		fmt.Println("You are underweight.")
	} else if bmi < 25 {
		fmt.Println("You have a normal/healthy weight.")
	} else if bmi < 30 {
		fmt.Println("You are overweight.")
	} else {
		fmt.Println("You are obese.")
	}
	
} 
