package main

import (
	"fmt"
	"sync"
	"time"
)

func numbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	defer wg.Done()
}

func alphabets(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go numbers(&wg)
	go alphabets(&wg)
	wg.Wait()
	fmt.Println("END MAIN")
}

/**
go rountine: thực hiện đồng thời. k cần chờ cho tới khi finish

---sync.WaitGroup(): dùng để chở cho tất cả các rountine finish rồi mới tiếp tục

Add(number): Thêm số lượng go rountine trong hàng chờ
Done(): Giảm 1 trong hàng chờ
Wait(): chờ cho tới khi hoàn thành hết rountines

---sync.Mutex: thực hiện lock code đảm bảo tại thời điểm duy nhất rountine thực hiện code section


*/
