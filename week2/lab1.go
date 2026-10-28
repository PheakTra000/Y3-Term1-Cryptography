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

	fmt.Println(num1 == num2)
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

}
