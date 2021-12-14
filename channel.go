package main

import "fmt"

func main() {
	//var a chan int
	//if a == nil {
	//	a = make(chan int)
	//	fmt.Printf("%T", a)
	//}

	a, b := 5, 10
	channel := make(chan int)
	go plus(a, b, channel)
	rs := <- channel
	fmt.Println(rs)

}

func plus(a, b int, channel chan int) {
	sum := a + b
	fmt.Println("sum in func: ", sum)
	channel <- sum
}
/*
	chn := make(chan int) -> tạo channel int
	chn1 := make(chan <- int) -> tạo channel chỉ để send(ghi)
	chn2 := make(<-chan int) -> tạo channel chỉ để đọc

	data := <- a // read from channel a
	a <- data // write to channel a

	=> channel: hiểu như là 1 cách nào đó để gửi nhận dữ liệu, nó communicate các go rountine

	channel có thể là 2 chiều gửi, nhận hoặc 1 chiều


	=================================================
	Buffered Channels and worker Pool

*/