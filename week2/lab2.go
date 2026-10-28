package main

import "fmt"

func main() {
	fmt.Println("Calculator")

	var num1 int
	var num2 int

	fmt.Print("Input num1: ")
	fmt.Scan(&num1)
	fmt.Print("Input num2: ")
	fmt.Scan(&num2)

	fmt.Println("num1 > 0 && num2 > 0: ", num1 > 0 && num2 > 0)

	fmt.Println("num1 > num2: ", num1 > num2)

	fmt.Println("num1 != num2: ", num1 != num2)
}
