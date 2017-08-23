package main

import (
	"fmt"
)

func operations(a, b int) (int, int, int) {
	var sum = a + b
	var minus = a - b
	var mul = a * b

	return sum, minus, mul
}

func main()  {
	a, b, c := operations(5, 6)

	fmt.Println("Addition =", a)
	fmt.Println("Substraction =", b)
	fmt.Println("Multiplication =", c)
}
