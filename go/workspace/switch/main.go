package main

import "fmt"

func checkSwitch() {
	// 1. Basic switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Case 1: Monday")
	case 2:
		fmt.Println("Case 1: Tuesday")
	case 3:
		fmt.Println("Case 1: Wednesday")
	default:
		fmt.Println("Case 1: Invalid day")
	}

	// 2. Multiple values in one case
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Case 2: Vowel")
	default:
		fmt.Println("Case 2: Consonant")
	}

	// 3. Switch without expression (acts like if-else)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Case 3: Grade A")
	case score >= 75:
		fmt.Println("Case 3: Grade B")
	case score >= 50:
		fmt.Println("Case 3: Grade C")
	default:
		fmt.Println("Case 3: Fail")
	}

	// 4. Type switch
	var x interface{} = 42
	switch v := x.(type) {
	case int:
		fmt.Println("Case 4: x is int with value", v)
	case string:
		fmt.Println("Case 4: x is string with value", v)
	default:
		fmt.Println("Case 4: Unknown type")
	}

	// 5. Fallthrough (forces next case to run)
	num := 2
	switch num {
	case 1:
		fmt.Println("Case 5: One")
		fallthrough
	case 2:
		fmt.Println("Case 5: Two (fallthrough from One if num==1)")
	case 3:
		fmt.Println("Case 5: Three")
	default:
		fmt.Println("Case 5: Other number")
	}
}

func main() {
	checkSwitch()
}
