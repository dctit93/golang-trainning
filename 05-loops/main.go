package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // nếu lớn hơn 5 thoát vòng lặp
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nKết thúc vòng lặp")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // nếu là số chẵn thì bỏ qua không chạy tiếp
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	/* Vòng lặp trong vòng lặp*/
	fmt.Println("Hai vòng lặp vòng nhau")
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
		}

	}
}
