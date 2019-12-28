package main

import (
	"fmt"
)
// Tiendc
/*
	Function có thể nhập vào biến đầu vào có số lượng không cố định


*/
func find(num int, nums ...int) {
	fmt.Printf("Kiểu giá trị của nums là %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "được tìm thấy tại vị trí thứ", i, "trong mảng", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "không được tìm thấy tại vị trí thứ", nums)
	}
	fmt.Printf("\n")
}
func findWithSlices(num int, nums []int) {
	fmt.Printf("Kiểu giá trị của nums là %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "được tìm thấy tại vị trí thứ", i, "trong mảng", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "không được tìm thấy tại vị trí thứ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	/*
		Theo khai báo của function trên thì  giá trị đầu tiên sẽ đc gán cho biến Num , những giá trị còn lại sẽ đc gán cho biến nums và nằm trong mảng có kiểu int
	*/
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	/*
		Trường hợp giá trị đầu vào của funciton là một slices với số phần tử ngẫu nhiên thì ta gọi function bằng cách ép kiểu sang slices
	*/
	findWithSlices(89, []int{89, 90, 95})

	nums := []int{89, 90, 95} // Nums là một slices trong khi hàm find giá trị đầu vào ko phải là slices thì ta gọi hàm bằng cách bên dưới
	find(89, nums...)

}
