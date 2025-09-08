package main

import "fmt"

func main() {
	// Integer types
	var age int = 30
	var score int32 = 95
	var bigNumber int64 = 1234567890

	// Floating-point types
	var pi float32 = 3.14
	var g float64 = 9.81

	// Boolean type
	var isGoFun bool = true

	// String type
	var name string = "Hareesh"

	// Default values (zero values)
	var defaultInt int       // 0
	var defaultBool bool     // false
	var defaultString string // ""

	// Short declaration (type inference)
	city := "Pune"
	year := 2025

	// Print values
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("Big Number:", bigNumber)
	fmt.Println("Pi:", pi)
	fmt.Println("Gravity:", g)
	fmt.Println("Is Go Fun?", isGoFun)
	fmt.Println("Name:", name)
	fmt.Println("Default int:", defaultInt)
	fmt.Println("Default bool:", defaultBool)
	fmt.Println("Default string:", defaultString)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)

}
