package main

import (
	"fmt"
)

/* Con trỏ

 */
func hello() *int {
	i := 5
	return &i
}
func main() {
	b := 255        // Giá trị của b là 255 và sẽ đc lưu tại địa chỉ bộ nhớ là 0xc0000a8000
	fmt.Println(&b) // 0xc0000a8000

	var a *int = &b // Kiểu a đc khai báo một biến con trò kiểu int nó trỏ tới địa chỉ vùng nhớ của b
	// a := &b  hoặc khai báo kiểu tự ép kiểu
	fmt.Printf("Kiểu giá trị của a là  %T\n", a)
	fmt.Printf("a là :  %v\n", a)
	fmt.Println("Giá trị tại ô nhớ của a", *a)

	/*
		Biến con trỏ được gán bằng địa chỉ ô nhớ
		Giá trị của biến con trỏ là giá trị của ô nhớ đó
		lấy giá trị của của biến con trỏ ta dùng * {tên biến}
		Vì thế muốn thay đổi giá tri thì ta gán *{tên biến }=  {giá trị}
	*/
	*a = 113
	fmt.Println("Giá trị tại ô nhớ của a sau khi thay đổi", *a)
	*a++
	fmt.Println("Giá trị tại ô nhớ của a sau khi thay đổi", *a)

	d := hello()
	fmt.Println("Value of d", *d)
}
