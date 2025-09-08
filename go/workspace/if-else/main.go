package main

import "fmt"

func checkConditions() {
	// 1. Simple if
	age := 20
	if age >= 18 {
		fmt.Println("Case 1: You are an adult.")
	}

	// 2. if … else
	if age < 18 {
		fmt.Println("Case 2: Minor")
	} else {
		fmt.Println("Case 2: Adult")
	}

	// 3. if … else if … else
	score := 75
	if score >= 90 {
		fmt.Println("Case 3: Grade A")
	} else if score >= 75 {
		fmt.Println("Case 3: Grade B")
	} else if score >= 50 {
		fmt.Println("Case 3: Grade C")
	} else {
		fmt.Println("Case 3: Fail")
	}

	// 4. if with short statement
	if num := 10; num%2 == 0 {
		fmt.Println("Case 4:", num, "is even")
	} else {
		fmt.Println("Case 4:", num, "is odd")
	}

	// 5. Nested if
	country := "India"
	if age >= 18 {
		if country == "India" {
			fmt.Println("Case 5: Eligible to vote in India")
		} else {
			fmt.Println("Case 5: Adult, but not in India")
		}
	} else {
		fmt.Println("Case 5: Not an adult yet")
	}

	// 6. if with logical operators
	hasID := true
	if age >= 18 && hasID {
		fmt.Println("Case 6: Entry allowed")
	} else if age >= 18 && !hasID {
		fmt.Println("Case 6: Age OK, but ID required")
	} else {
		fmt.Println("Case 6: Not allowed")
	}
}

func main() {
	checkConditions()
}
