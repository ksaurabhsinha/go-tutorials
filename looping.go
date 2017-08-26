package main

import (
	"fmt"
)

func main() {

	//Go has only FOR loop ;)

	//Print from 0 to 99
	for i := 0; i < 100; i++ {
		fmt.Println("Number is =", i)
	}

	//Optional Init and Post Statements in for loop
	sum := 1
	for ; sum < 100; {
		sum += sum
	}

	fmt.Println(sum)

}
