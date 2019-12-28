package main

import (
	"fmt"
)

func ChangeInsideFunction(num [5]int) {
	num[0] = 100
	fmt.Println("inside function ", num)

}
func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func main() {
	var a [3]int // khai báo mảng a kiểu int và có 3 phần tử
	a[0] = 12    // gán giá trị cho các phần tử của mảng
	a[1] = 78
	a[2] = 50
	fmt.Println(a) // [12 78 50]

	/*
		Chúng ta có thể khai báo giá trị ban đâu bằng cách sau
	*/
	b := [3]int{12, 78, 50}
	fmt.Println(b) // [12 78 50]

	/*
		Trường hợp khai báo giá trị ban đầu nhưng không đủ các phần tử của mảng
	*/
	c := [5]int{12}
	fmt.Println(c) // Các giá trị phần tử còn lại sẽ đc gán giá trị mặc định giống như khai báo biến không có giá trị, giá trị in ra sẽ là [12 0 0 0 0]

	/*
		Khai báo kiểu array không cần quan tâm đến độ dài mảng, go sẽ tự map theo số phần tử
	*/
	d := [...]string{"USA", "China", "India", "Germany", "France", "VietNam"}
	fmt.Println("Độ dài mảng là:", len(d))

	/*
		Các trường hợp đặc biệt
	*/
	num := [...]int{5, 6, 7, 8, 9}
	fmt.Println("Trước khi qua function", num)
	ChangeInsideFunction(num)
	fmt.Println("Sau khi qua function ", num) // Sau khi qua function thì giá trị của num không đổi , nó chỉ là một giá trị đầu vào của function và ko ảnh hưởng khi nằm ngoài function
	/*
		Thao tác với giá trị của mảng
	*/
	for i := 0; i < len(num); i++ {
		fmt.Printf(" Phần tử %d giá trị là  %v\n", i, num[i])
	}
	// Chúng ta có thể viết ngăn gọn hơn bằng cách
	fmt.Println("Kết quả chạy với range")
	for i, v := range num { //range returns both the index and value
		fmt.Printf(" Phần tử %d giá trị là  %v\n", i, v)
	}
	// Hoặc
	fmt.Println("Trả kết quả mà ko cần index")
	for _, v := range num {
		fmt.Printf("giá trị là  %v\n", v)
	}

	/*
		MẢNG HAI CHIỀU
	*/
	e := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(e)
	var f [3][2]string
	f[0][0] = "apple"
	f[0][1] = "samsung"
	f[1][0] = "microsoft"
	f[1][1] = "google"
	f[2][0] = "AT&T"
	f[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(f)

	/*
		=================================================== Slices ======================================================
		Slices giống như một công cụ giúp ta thao tác với các đoạn của mảng một cách thuận tiện nhất
		Slies cũng có thề là một array nhưng khai báo số thứ tự của mảng là nil
	*/
	g := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Trước khi qua thao tác với slices: ", g) //1 2 3 4 5

	nums1 := g[:] // khai báo với tất cả phần tử của mảng
	fmt.Println(nums1)
	//var slices []int = g[1:4]
	//fmt.Println(slices)
	// hoặc có thể viết

	gslices := g[1:4]    // lấy giá trị từ 2 -> 4 , mảng bắt đầu từ 0
	fmt.Println(gslices) // 2 3 4

	for i := range gslices {
		fmt.Println("i", i)
		gslices[i]++ // cộng lên 1 cho mỗi phần tử của slice
	}
	fmt.Println("Sau khi qua slices", g)

	// Sau khi thông qua sử lý với silce thì giá trị của mảng ban đầu sẽ thay đổi khi slice của nó thay đổi

	/*
		Khai báo một clices
	*/
	cars := []string{"Ferrari", "Honda", "Ford"}                                       // cars là một slices
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3

	var namesSlices []string // Nếu mà khai báo slices ko có giá trị thì slices đó giá trị mặc định là nil
	if namesSlices == nil {
		fmt.Println("slice is nil going to append")
		namesSlices = append(namesSlices, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", namesSlices)
	}

	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos) // khác với array một slice khi qua một function thì giá trị của clices đó sẽ thay đổi
}
