package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			s := fmt.Sprintf("putting x = %d in channel y = %d", x, y)
			fmt.Println(s)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			s := fmt.Sprintf("%d = %d", i, <-c)
			fmt.Println(s)
			time.Sleep(100 * time.Millisecond)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
