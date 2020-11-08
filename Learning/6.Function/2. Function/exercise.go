package main

import (
	"fmt"
	"strconv"
)

//function and parameter

//creating getBiography function, to show someone's identity
//showing multiple return on one function
//declaring returned variable in a function
func getBiography(age int, name string, status string) (bio string, ageIn10 int) { //this function will return one string and one integer

	//converting age to string so that it can be used in a sentence
	ageNow := strconv.Itoa(age)

	//first variable that will be returned, showing someone's biography
	bio = name + " is a " + status + ", right now he is " + ageNow + " years old. "

	//second variable that will be returned, predict age in 10 years
	ageIn10 = age + 10

	//returning two variable at once
	return
}

func main() {

	//using and printing getBiography function

	//splitting one function into two variable
	Info, agePrediction := getBiography(60, "Bill Gates", "good people") //converting getBiography function into two strings

	//printing two result
	fmt.Println(Info)
	fmt.Println("his age in 10 years is ", agePrediction)
}
