package main

import (
	"03-packages/hinhchunhat" // Đây là package customer
	"03-packages/hinhvuong"
	"fmt" // Đây là package của go
)

// Để import package khác vào đây để sử dụng thì ta có hai  cách
// Chép source 03-packages đặt trong thư mục GOPATH mà bạn đã cài đặt - cách này không khuyến khích sử dụng vì một project ta sử dụng rất nhiều package bên thứ 3 , và điều này khá là ràng buộc
// cách 2 là sử dụng go module: khuyến khích sử dụng
// sử dụng go mod init {tên project}
// Nó sẽ tạo một thư file go.mod và muốn sử dụng package nào của mình thì chỉ cần import từ thư mục cha của project
func init() {
	fmt.Println("Đây là hàm main")
}
func main() {
	/* Chú ý
	Khi sử dụng hàm init tại package thì khi import packege thì hàm đó sẽ đc chạy trước
	Vậy giá trị in ra sẽ là :

		Đây là package hinhchunhat
		Đây là package hinhvuong
		main
		Chu vi hinh chữ nhật là: 12
		Chu vi hình vuông là: 16

	*/
	chuvihinhchunhat := hinhchunhat.ChuVihinhChuNhat(3, 3)
	//chuvihinhchunhat := hinhchunhat.chuVihinhChuNhat(3, 3)
	/* Chú ý
	Khi khai báo một function trong một package thì khi đặt tên function đó thì phải Viết hoa chữ cái đầu
	*/
	fmt.Println("Chu vi hinh chữ nhật là:", chuvihinhchunhat)
	chuvihinhvuong := hinhvuong.ChuViHinhVuong(4)
	fmt.Println("Chu vi hình vuông là:", chuvihinhvuong)
}
