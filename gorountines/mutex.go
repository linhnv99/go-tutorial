package main

import (
	"fmt"
	"sync"
)

var (
	count = 0
	wg = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(2)

		mu.Lock() // lock lại cho tới khi in ra
		go printCount()
		wg.Done()

		mu.Lock()
		go increment()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("END")
}

func printCount() {
	fmt.Println(count)
	mu.Unlock()// mở sau khi in success
}

func increment() {
	count++
	mu.Unlock()
}
