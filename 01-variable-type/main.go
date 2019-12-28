// TIENDC
package main

import (
	"fmt"
)

func main() {
	/*  ============================================= TYPE =============================================
	Chúng ta có các kiểu dữ liệu cơ bản như sau
	- bool: Kiểu dữ liệu trả về true hoặc false
	- Numberic:
		+ int8, int16, int32, int64, int :  kiểu số nguyên có tồn tại giá trị âm
			* int8 : -128 – 127
			* int16 : -32768 – 32767
			* -2147483648 – 2147483647
			* -9223372036854775808 – 9223372036854775807
		+ uint8, uint16, uint32, uint64, uint
			* uint8 : 0 – 255
			* uint16 : 0 – 65535
			* uint32 : 0 – 4294967295
			* uint64 : 0 – 18446744073709551615
		+ float32, float64 : Số thực có giá trị thập phân
		+ complex64, complex128
		+ byte
		+ rune
	*/

	/* ============================================= VARIABLES ===========================================
	Chúng ta có các cách khai báo biến như sau :
	+ Cách 1
		var {Tên biến } {Kiểu giá trị} = {giá tri}

	*/
	// Trường hợp khai báo bình thường
	var a int = 10
	fmt.Println("Giá trị của a", a)
	// Khai báo không gán giá trị thì biến đó sẽ được GO gán giá trị mặc định
	// Kiểu int thì là giá trị 0
	var age int
	fmt.Println("Tuổi của tôi", age)
	age = 18
	fmt.Println("age", age)
	// Kiểu string là ""
	var name string
	fmt.Println("Tên của tôi", name)
	name = "Tèo"
	fmt.Println("name", name)
	// Mặc định kiểu bool là false
	var value bool
	fmt.Println("Value is", value)
	value = true
	fmt.Println("value", value)
	// ép kiểu
	var b = 19
	fmt.Printf("Kiểu giá trị là %T : %v \n", b, b) //  Trong đó %T là kiểu giá tri , %d là số nguyên , %s kiểu string , %f kiểu float ,%v là giá trị nào cũng đc ^^
	/*
		+ Cách 2
			{tên biến } := {giá trị}
	*/
	// Cách khai báo ngắn gọn nhất
	c := 20
	fmt.Printf("Kiểu giá trị là %T : %d \n", c, c)
	d := "Đỗ Công Tiền"
	fmt.Printf("Kiểu giá trị là %T : %s \n", d, d)
	fmt.Println("Tổng của b và c: ", b+c)
	tong := b + c
	fmt.Println("tong", tong)
	/*
		Khai báo với const
	*/
	const constValue = 55
	//constValue = 89  giá trị của biến const không thể được gán lại bằng một giá trị khác
	//const constValue = math.Sqrt(4)  Giá trị của biến const không đước khai báo bằng cách gán giá trị trả về của một function

}
