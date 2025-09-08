package main

import "fmt"

func main() {
	// 1. Declare an array with fixed size
	var arr1 [5]int
	fmt.Println("Case 1: Empty array of 5 ints:", arr1)

	// 2. Initialize with values
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Case 2: Initialized array:", arr2)

	// 3. Initialize without specifying size (compiler infers size)
	arr3 := [...]string{"Go", "Java", "Python"}
	fmt.Println("Case 3: Inferred size array:", arr3)

	// 4. Access elements by index
	fmt.Println("Case 4: First element of arr2:", arr2[0])
	fmt.Println("Case 4: Last element of arr2:", arr2[4])

	// 5. Modify elements
	arr2[2] = 99
	fmt.Println("Case 5: Modified arr2:", arr2)

	// 6. Loop through array using index
	fmt.Println("Case 6: Loop with index")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Index %d -> %d\n", i, arr2[i])
	}

	// 7. Loop with range
	fmt.Println("Case 7: Loop with range")
	for index, value := range arr3 {
		fmt.Printf("Index %d -> %s\n", index, value)
	}

	// 8. Multidimensional array
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Case 8: 2x3 Matrix:", matrix)

	// Loop through matrix
	fmt.Println("Case 8: Matrix elements:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	check()
}

func check() {
	// 1. Declare and initialize arrays
	var arr1 [5]int                           // default values (0)
	arr2 := [5]int{10, 20, 30, 40, 50}        // initialized
	arr3 := [...]string{"Go", "Java", "Rust"} // inferred size

	fmt.Println("Case 1: arr1 =", arr1)
	fmt.Println("Case 1: arr2 =", arr2)
	fmt.Println("Case 1: arr3 =", arr3)

	// 2. Access & modify elements
	fmt.Println("Case 2: arr2[0] =", arr2[0]) // first element
	arr2[2] = 99                              // modify element
	fmt.Println("Case 2: arr2 after update =", arr2)

	// 3. Length of array
	fmt.Println("Case 3: len(arr2) =", len(arr2))

	// 4. Classic for loop
	fmt.Println("Case 4: Classic for loop")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Index %d -> %d\n", i, arr2[i])
	}

	// 5. While-style loop
	fmt.Println("Case 5: While-style loop")
	j := 0
	for j < len(arr3) {
		fmt.Printf("Index %d -> %s\n", j, arr3[j])
		j++
	}

	// 6. Infinite loop with break
	fmt.Println("Case 6: Infinite loop with break")
	k := 1
	for {
		if k > 5 {
			break
		}
		fmt.Print(k, " ")
		k++
	}
	fmt.Println()

	// 7. Loop with continue (skip even numbers)
	fmt.Println("Case 7: Using continue")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 8. Range loop over array
	fmt.Println("Case 8: Range loop over arr3")
	for index, value := range arr3 {
		fmt.Printf("Index %d -> %s\n", index, value)
	}

	// 9. Range loop ignoring index
	fmt.Println("Case 9: Range loop ignoring index")
	for _, value := range arr2 {
		fmt.Println("Value:", value)
	}

	// 10. Nested loops (multiplication table)
	fmt.Println("Case 10: Nested loops (3x3 table)")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}

	// 11. Multidimensional array
	fmt.Println("Case 11: Multidimensional array")
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
