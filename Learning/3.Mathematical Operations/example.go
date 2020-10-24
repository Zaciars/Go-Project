package main

import (
	 "fmt"
)

func main() {

	var result int
	var result2 int
	var result3 int

	a := 10
	b := 20
	c := 5
	d := 12
	e := 8

//times = "*",
//divide = "/",
//modulo(remainder after dividing) = %,
//plus/minus one(+/-1) = (var)++/--

	result = a + b - c
	result2 = a * b / c
	result3 = result % result2
	d++
	e--

	fmt.Println("the result is", result)

	fmt.Println("the result is", result2)

	fmt.Println("the result is", result3)

	fmt.Println("the result is",d)

	fmt.Println("the result is",e)

}
