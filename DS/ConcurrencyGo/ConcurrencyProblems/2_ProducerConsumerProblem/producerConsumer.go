package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go Producer(ch)
	Consumer(ch)
}

func Producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Procuded: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
func Consumer(ch <-chan int) {
	for data := range ch {
		fmt.Printf("Consumed:%d\n", data)
		time.Sleep(1 * time.Second)
	}
}
