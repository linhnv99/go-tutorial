package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "Two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("5 seconds")
		}
	}
}

/**
2 loại: Buffered và unbuffered

select có thể giúp thực hiện chọn từ nhiều channel != nhau
*/
