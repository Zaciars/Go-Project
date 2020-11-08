package main

import (
	"fmt"
)

func main() {

	debt := 20000
	money := 10000


	//if you want to test more than one case
	if money > debt { //case one
		fmt.Println("Happy")
	} else if money == debt { //case two
		fmt.Println("a bit sad")
	}else { //fallback if the result doesn't pass the test
		fmt.Println("depression")
	}

}
