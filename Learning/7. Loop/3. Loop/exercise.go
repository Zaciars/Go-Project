package main

import (
	"fmt"
)

func main() {

	//n + (n+1) + ...
	//we want to count  this : 1 + 2 + 3 + ... + 10 = x

	x := 0

	for num := 1 ; num <= 10 ; num++ {
		x = x + num
	}

	/*
	How this work :
	first we declare a variable called "x" as 0 (starting point)
	then, we create a loop where "x" variable added by number 1 until 10(num variable)
	 */


	fmt.Println("amount of series of number up to 10", x)
}
