package main

import (
	"fmt"
)

//function
//this is function returninng 1 integer
func printNumber()  int{

	fmt.Println("you called the printNumber function")

	//what value will be returned?
	//declarating retured value
	return 10

}
	//creating parameter ( value that we pass over)
	//for example this is function with two parameter
func multiply(number1 int, number2 int)  int{

	fmt.Println("you called the multiply function")

	return number1 * number2

}

func main() {

	fmt.Println("this function return", printNumber())

	fmt.Println("this function return", multiply(5, 10))

}
