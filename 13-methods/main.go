package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
	Phương thức (Method) là một hàm (function) được khai báo cho riêng một kiểu dữ liệu đặc biệt, kiểu dữ liệu này được gọi là vật nhận (receiver) nó được đặt giữa từ khóa func và tên của phương thức.
	Vật nhận này có thể là kiểu struct (cấu trúc) hoặc non-struct (phi cấu trúc).
	Vật nhận phải có sẵn để truy cập bên trong phương thức.
	Ví dụ dưới đây Employee là một kiểu dữ liệu đặc biệt
*/

//method
func (e Employee) displaySalary(heSo int) int {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
	return e.salary * heSo
}

// function

func displaySalaryFunc(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// Trường hợp đặc biệt
type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	tongLuong := emp1.displaySalary(3)
	fmt.Println(tongLuong)
	displaySalaryFunc(emp1)

	/*
		Phương thức fullAddress đc truy cập bởi kiểu address nhưng address lại đc chứa trong kiểu person nên person có toàn quyền truy cập vào fullAddress
	*/
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress()

}
