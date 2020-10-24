package main

import (
	"fmt"
	"strconv"
)

func main() {

	debt := 5000

	//converting string to in using Atoi, by creating a new variable incase of error
	salary := "2000"
	salary2,_ := strconv.Atoi(salary)

	bonus := salary2 + 1500


	//convert int to string using Itoa
	fmt.Println("your debt is " + strconv.Itoa(debt))

	fmt.Println("your salary is", bonus)

}
