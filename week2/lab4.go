package main

import (
	"fmt"
)

func main() {

	menu()

	// init variable
	var choice int = userChoice()
	num1, num2 := userInput()

	// method
	switch choice {
	case 1:
		fmt.Println("Result: ", add(num1, num2))
	case 2:
		fmt.Println("Result: ", sub(num1, num2))
	case 3:
		fmt.Println("Result: ", mul(num1, num2))
	case 4:
		fmt.Println("Result: ", div(num1, num2))
	case 5:
		fmt.Println("Result: ", mod(num1, num2))
	case 6:
		fmt.Println("Exit!")
	}

}

func userChoice() int {
	var choice int
	fmt.Print("Choose: ")
	fmt.Scan(&choice)
	return choice
}

func userInput() (int, int) {
	var num1 int
	var num2 int

	fmt.Print("Enter num1: ")
	fmt.Scan(&num1)

	fmt.Print("Enter num2: ")
	fmt.Scan(&num2)

	return num1, num2
}

func add(a int, b int) int {
	return a + b

}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) any {
	// return a / b
	if b == 0 {
		return "Division by zero"
	} else {
		return a / b
	}
}

func mod(a int, b int) int {
	return a % b
}

func menu() {
	fmt.Println("=====[Mini Calculator]=====")
	fmt.Println("1. Add")
	fmt.Println("2. Sub")
	fmt.Println("3. Mul")
	fmt.Println("4. Div")
	fmt.Println("5. Mod")
	fmt.Println("6. Exit")
}
