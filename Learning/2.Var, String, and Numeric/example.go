package main

//this is how to comments

import (
	"fmt"
)

//variable - tipe data = string

var  study = "Go-lang "
var number = "1 day!!!"

//combining two variable

var name1 = "Alif"
var name2 = "Zakiansyah"
var (
	nameFull = name1 + " " + name2
)

//declaring two variable in one line

var date,place = "2006","Palembang"

//declaring empty variable
var blank string
var null int

//printing the results

func main() {

	//shorthand variable (in function only)
	example := "fast variable declaration test"

	//variable dataType	:	numeric
	//numericType		:	int/integer
	age := 14

	fmt.Println("hello world, i have studied " + study + "for" + number + "")

	fmt.Println("my name is " + nameFull)

	fmt.Println("you can call me " + name1)

	//printing int
	fmt.Println("i was ",age ,"years old when i wrote this")

	fmt.Println("I was bon in " + place + "," + date )

	fmt.Println("this should've been empty : " + blank)

	fmt.Println(example, null)

	//adding int
	a := 20
	b := 10
	var total = a - b

	fmt.Println("20 + 10 equals to",a + b)

	fmt.Println("20 - 10 equals to",total)
}
