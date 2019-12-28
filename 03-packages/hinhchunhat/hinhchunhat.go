package hinhchunhat

import (
	"fmt"
)

func init() {
	fmt.Println("Đây là package hinhchunhat")
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
