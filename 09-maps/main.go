package main

import (
	"fmt"
)

/*
	Map quản lý phần tử theo key và value
	Khai báo
	+ Khi khai báo một map không được cho giá trị nil , nên khi khai báo mà không có giá trị sẽ báo lỗi, nếu khai báo không có giá trị thì dùng make
		make(map[{kiểu giá trị của key}]{kiểu giá trị của value})
*/
func main() {
	// Khai báo với map rỗng
	var personSalary = make(map[string]int)

	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000

	fmt.Println("Map contents:", personSalary)
	// Khai báo với map có giá trị đầu vào
	personSalaryWithValue := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("personSalaryWithValue map contents:", personSalaryWithValue)

	// Lấy giá trị từ map
	fmt.Println("Salary of joe is", personSalaryWithValue["steve"])

	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalaryWithValue[employee])

	// Khi truy xuất giá trị của một map ta có cờ như sau

	newEmp := "joe"
	value, ok := personSalaryWithValue[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	// Lấy tất cả giá trị của một map

	for key, value := range personSalaryWithValue {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	// Xóa giá trị trong một map
	fmt.Println("map before deletion", personSalaryWithValue)
	delete(personSalaryWithValue, "jamie")
	fmt.Println("map after deletion", personSalaryWithValue)

}
