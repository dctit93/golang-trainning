// Tiendc
package main

import (
	"fmt"
)

func main() {
	/*
		if else của Go thì không khác gì các ngôn ngữ lập trình khác
	*/
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

	b := false

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	// true false
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
	// not
	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}
	// or 
	if true || false {
		fmt.Println("This ran")
	}
	// and
	
	if true && false {
		fmt.Println("This did not run")
	}
	
}
