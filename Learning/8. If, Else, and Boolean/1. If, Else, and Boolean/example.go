package main

import (
	"fmt"
)

func main() {

	//if and else
	color := "green"
	num   := 10

	//type data boolean(true/false)
	if true {
		println("true mint") // blank var will always be true
	}

	//for string
	//if the color is green, "if" statement will be used.
	if color == "green" {
		fmt.Println("you may pass")
	} else {
		fmt.Println("stop") //fallback if "color" variable is not green

	}

	//for integer
	//if the "num" value is bigger than 100, will print the statement
	//if "num" is lower/equal to 100, will do nothing
	if num > 100 {
		fmt.Println("the number is bigger than 100") //"if" statement can be used without "else"
	}

}
