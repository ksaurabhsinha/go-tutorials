package main

import (
	"fmt"
)

//Constants can only be declared outside the function
const Pi = 3.14


func main() {

	//Type interface
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	//Type Conversions
	a := 42
	b := float64(a)
	c := uint(b)

	fmt.Println("Value of Pi =", Pi)

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)

	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
