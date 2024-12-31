package main

import (
	"fmt"
	"sync"
)

// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
/*
Ping:1
Pong:1
Ping:2
Pong:2
Ping:3
Pong:3
Ping:4
Pong:4
Ping:5
Pong:5
Ping:6
Pong:6
Ping:7
Pong:7
Ping:8
Pong:8
Ping:9
Pong:9
Ping:10
Pong:10
*/
func main() {
	maxPingPongLen := 10
	pingChan := make(chan int)
	pongChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= maxPingPongLen; i++ {
			fmt.Printf("Ping:%d\n", i)
			pingChan <- i
			<-pongChan
		}
		close(pingChan)
	}()
	go func() {
		for k := range pingChan {
			fmt.Printf("Pong:%d\n", k)
			pongChan <- k
		}
		close(pongChan)
	}()

	wg.Wait()
}

/*
package main

import (
	"fmt"
	"sync"
)
func main() {
	maxPingPongLen := 10
	pingChan := make(chan int)
	pongChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(maxPingPongLen)
	for i := 1; i <= maxPingPongLen; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Ping:%d\n", i)
			pingChan <- i
			<-pongChan
		}(i)
	}

	go func() {
		for k := range pingChan {
			fmt.Printf("Pong:%d\n", k)
			pongChan <- k
		}
		close(pongChan)
	}()

	wg.Wait()
	close(pingChan)
}
*/
