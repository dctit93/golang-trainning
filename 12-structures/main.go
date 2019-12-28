package main

import (
	"fmt"
)

// Khai báo một struct HocSinh với các thành phần hoTen, ngaySinh,diaChi,tuoi
type HocSinh struct {
	hoTen, ngaySinh, diaChi string
	tuoi                    int
}

func main() {
	// Tạo một học sinh
	hocSinh1 := HocSinh{
		hoTen:    "Đỗ Công Tiền",
		ngaySinh: "1/1/2001",
		tuoi:     18,
		diaChi:   "Bến tre",
	}

	//creating structure without using field names
	hocSinh2 := HocSinh{"Nguyễn Thị Tèo", "1/1/2019", "Long An", 18}

	fmt.Println("Học sinh 1", hocSinh1)
	fmt.Println("Học sinh 2", hocSinh2)

	// Gán giá trị trực tiếp
	hocSinh3 := struct {
		hoTen, ngaySinh, diaChi string
		tuoi                    int
	}{
		hoTen:    "Trần Văn Tí",
		ngaySinh: "1/1/2001",
		tuoi:     18,
		diaChi:   "Bến tre",
	}

	fmt.Println("Học sinh 3", hocSinh3)

	// Trường hợp khai báo mà không gắn giá trị thì từng trường của struct sẽ có giá trị mặc định

	var hocSinh4 HocSinh
	fmt.Println("Học sinh 4", hocSinh4) //{   0}

	// Tương tự nếu ta gán chĩ một trường trong struct, các giá trị khác ko gán giá trị thì nó cũng sẽ đc gán giá trị mặc định

	hocSinh5 := HocSinh{
		hoTen: "Phạm Văn Mách"}
	fmt.Println("Học sinh 5", hocSinh5) //{Phạm Văn Mách   0}

	// Lấy giá trị từ struct

	fmt.Println("Tên của Học sinh 5", hocSinh5.hoTen) //Phạm Văn Mách

	// Tương tự ta có thể gán lại giá trị khác với hoTen

	hocSinh5.hoTen = "JackMa"

	fmt.Println("Tên của Học sinh 5 sau khi thay đổi", hocSinh5.hoTen) //JackMa

}
