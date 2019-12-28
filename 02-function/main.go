//TienDC
package main

import (
	"fmt"
)

/*
	Cách khai báo funciton:
	func {tên function} ( {giá tri đầu vào} {Kiểu giá trị đầu vào},.....) {kiểu dữ liệu trà về} {
		function body
	}
	chú ý:
	Tên của funciton phải viết hoa- chuyện này không bắt buộc, giả sử func nằm trong cùng một package thì ta có thể sử dụng bình thường nhưng nếu sử dụng cho một package khác thì khi import vào thì GO sẽ không tìm thấy
*/
func CalculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func HinhChuNhat(dai, rong int) (int, int) {
	var chuVi = (dai + rong) * 2
	var dienTich = dai * rong
	return chuVi, dienTich
}
func HinhChuNhatNameReturn(dai, rong int) (chuVi int, dienTich int) {
	chuVi = (dai + rong) * 2
	dienTich = dai * rong
	return chuVi, dienTich
}

func main() {
	price, no := 90, 6
	totalPrice := CalculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	dai, rong := 2, 7
	//chuVi, dienTich := HinhChuNhat(dai, rong)
	chuVi, dienTich := HinhChuNhatNameReturn(dai, rong)
	//chuVi, _ := HinhChuNhatNameReturn(dai, rong)
	fmt.Printf("Chu vi : %v và diện tích là %v\n", chuVi, dienTich)
	//fmt.Printf("Chu vi : %v \n", chuVi)
}
