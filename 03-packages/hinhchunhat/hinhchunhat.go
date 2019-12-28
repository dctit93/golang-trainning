package hinhchunhat

import (
	"fmt"
)
// Ta thấy dù biến bien nằm ngoài file hinhchunhat.go nhưng cùng thuộc package hinhchunhat nên ta sử dụng được
func init() {
	fmt.Println("Đây là package hinhchunhat")
	fmt.Println(bien)
}
func ChuVihinhChuNhat(dai int, rong int) int {
	
	return (dai + rong) * 2
}

func chuVihinhChuNhat(dai int, rong int) int {
	return (dai + rong) * 2
}

func DienTichHinhChuNhat(dai int, rong int) int {
	return dai * rong
}
