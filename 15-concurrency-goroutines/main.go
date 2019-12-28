package main

/*
	Concurrency là khả năng thực hiện nhiều tác vụ cùng một lúc.
	Ví dụ như lúc đang ăn ta có thể vừa ăn vừa bấm điện thoại chẳng hạn

	Goroutines là các hàm hoặc phương thức chạy đồng thời với các hàm/ phương thức khác.
	Goroutines có thể được coi là những luồng gọn nhẹ. Chi phí tạo một Goroutine tương đối thấp so với một luồng.
	Do vậy, những ứng dụng Go có hàng ngàn Goroutines chạy đồng thời là điều hết sức bình thường.

	Goroutines trao đổi thông qua các kênh. Các kênh được thiết kế để ngăn ngừa các khả năng xung đột xảy ra khi truy cập bộ nhớ chia sẻ sử dụng Goroutines.

	Nhập từ khóa go vào trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời.
*/
import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	/*
		Khi khởi chạy hàm main thì ta sẽ có:

		Thực chất hàm main cũng là một Goroutine, khi chạy thì hàm main sẽ chạy đồng thời với hàm hello ,nhưng hàm main là một Goroutine chính còn hàm hello là một hàm
		Goroutine phụ.

		Trong gọi lệnh go hello() thì tính năng Goroutine sẽ không chờ hàm hello chạy xong mà in ra dòng code kế tiếp ngay.
		Sau đó Goroutine chính dừng nên hello Goroutine không còn cơ hội để chạy.


	*/
	go hello()
	fmt.Println("main function")
	//
	/*
		ở ví dụ bên dưới Goroutine chính khi chạy sẽ gọi hàm numbers nhưng không chờ numbers chạy xong chương trình sẽ gọi liền alphabets, tương tự thế chương trình cũng
		không chờ alphabets chạy xong sẽ lập tức gọi thời gian ngủ của goroutine chính là 3s
		Trong thời gian ngủ của gorountine chính thì các goroutine phụ sẽ có cơ hội chạy ( chạy đồng thời)
	*/
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
