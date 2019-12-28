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
}
