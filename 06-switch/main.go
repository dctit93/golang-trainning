package main

import (
	"fmt"
)

func main() {
	finger := 10
	switch finger {
	case 1:
		fmt.Println("Số một")
	case 2:
		fmt.Println("Số hai")
	case 3:
		fmt.Println("Số ba")
	case 4:
		fmt.Println("Số bốn")
	case 5:
		fmt.Println("Số năm")
	case 6, 7, 8, 9, 10, 11: // Khai báo nhiều giá trị trong một case
		fmt.Println("Số từ 6 đến 11")
	default: //default case
		fmt.Println("Số lạ")
	}
	/*
		Check case đúng sai
	*/
	num := 75
	switch { // no expression
	case num >= 0 && num <= 50: // Check true false ở đây
		fmt.Println("Từ 0 đến 50")
	case num >= 51 && num <= 100:
		fmt.Println("Từ 51 đến 100")
	case num >= 101:
		fmt.Println("Lớn hơn 101")
	}
}
