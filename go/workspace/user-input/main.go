package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")

	// %s for string, %d for integer
	fmt.Scanf("%s %d", &name, &age)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
