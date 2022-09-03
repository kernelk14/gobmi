package main

import (
	"fmt"
)

// This is a BMI Counter in Golang

var printf = fmt.Printf
var print = fmt.Print
var HEIGHT float64
var WEIGHT int

// TODO: Add Indicators for underweight, normal, and obese.
// TODO: Save details in form of a text file.
// TODO: Fix the indicator for overweight.
func main() {
	print("Enter your height (in meters): ")
	fmt.Scanln(&HEIGHT)
	print("Enter your weight (in kilograms): ")
	fmt.Scanln(&WEIGHT)

	var total_height = HEIGHT * HEIGHT
	var total_bmi = float64(WEIGHT) / total_height

	printf("Your BMI is: %.1f\n", total_bmi)

	// Now, for the moment of truth
	// For Underweight
	if total_bmi < 18.5 {
		print("FINDINGS: You are underweight, please consult your doctor.\n")
	}
	// For Normal
	else if total_bmi >= 18.5 {
		if total_bmi <= 29.9 {
			print("FINDINGS: You are in a normal state, continue live healthy.\n")
		}
	} 
	// For Overweight
	else if total_bmi >= 30.0 {
		if total_bmi > 30.0 {
			print("FINDINGS: You are overweight/obese, please consult your doctor.\n")
		}
	} 
	// If something weird happens
	else {
		print("Unhandled exception: %.1f: Cannot interpret BMI count.\n", total_bmi)
	}
}
