package main

import "fmt"

func calculations(a, b int) (sum, sub, mul int)  {
	sum = a + b
	sub = a - b
	mul = a * b

	return
}

func main()  {
	a, b, c := calculations(5, 6)

	fmt.Println("Addition =", a)
	fmt.Println("Substraction =", b)
	fmt.Println("Multiplication =", c)
}
