package main

import "fmt"

var globalVar string = "Global Variable"

const PI float64 = 3.14 // global constant

func main() {
	fmt.Println(globalVar)

	var number1 int = 1
	var number2 uint8 = 2
	fmt.Println(number1)
	fmt.Println(number2)

	var number3 uint64 = 923607624502350325
	var number4 int64 = -923607624502350325
	fmt.Println(number3)
	fmt.Println(number4)

	check := "Harsh"
	check = "Harry"
	fmt.Println((check))

	{
		hey := "Many"
		fmt.Println((hey))
	}

	var (
		name     string = "Hmm"
		checking int    = 3
	)

	fmt.Println(name)
	fmt.Println(checking)

	go example()

}

func example() {
	var radius float64 = 5.154
	var area float64

	area = PI * radius * radius
	fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
	fmt.Printf("Area of Circle is : %f", area)
	fmt.Println("Thank You")
}
