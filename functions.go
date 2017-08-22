package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func divide(a int, b int) int {
	return a / b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Addition =", add(5, 6))
	fmt.Println("Substraction =", subtract(5, 6))
	fmt.Println("Division =", divide(5, 6))
	fmt.Println("Multiplication =", multiply(5, 6))
}
