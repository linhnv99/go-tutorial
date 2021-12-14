package main

import (
	"fmt"
	"sync"
)
var rs string
func main() {

	// input: 10 file
	// tạo 10 luồng để đọc 10 file đấy

	links := []string{"a", "b", "c"}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, v := range links {
		wg.Add(1)
		go readFile(v, &wg, &mu)
	}

	wg.Wait()
	fmt.Println(rs)
}

func readFile(file string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	rs += file
	mu.Unlock()
	fmt.Println(rs)
}
