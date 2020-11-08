package main

import (
	"fmt"
)

func payDebt(money int) (text string) {

	debt := 50000

	if money > debt {
		text = "Happy"
	} else if money == debt {
		text = "a bit sad"
	} else {
		text = "depression"
	}

	return text
}

func main() {
	fmt.Println(payDebt(0))
}
