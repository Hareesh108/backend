package main

import "fmt"

func checkLoops() {
	// 1. Classic for loop (like C/Java)
	fmt.Println("Case 1: Classic for loop")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. While-style loop (only condition)
	fmt.Println("Case 2: While-style loop")
	j := 1
	for j <= 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// 3. Infinite loop (break required)
	fmt.Println("Case 3: Infinite loop with break")
	k := 1
	for {
		if k > 5 {
			break
		}
		fmt.Print(k, " ")
		k++
	}
	fmt.Println()

	// 4. Using continue (skip even numbers)
	fmt.Println("Case 4: Using continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5. Range loop (over slice)
	fmt.Println("Case 5: Range over slice")
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 6. Range loop (over string)
	fmt.Println("Case 6: Range over string")
	word := "GoLang"
	for index, char := range word {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// 7. Nested for loops
	fmt.Println("Case 7: Nested loops (multiplication table 1-3)")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	checkLoops()
}
