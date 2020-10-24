package main

import (
	"fmt"
)

//Const
func main() {
	var a = 10
	a++

	const debt = 50
	//debt--
	//debt = 10
	//constant value cannot be changed


	fmt.Println("the result is", a)

	fmt.Println("your debt is", debt)
}
