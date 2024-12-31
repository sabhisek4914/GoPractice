// CODE EXAMPLE VALID FOR COMPILING
package main

import (
	"fmt"
	"sync"
)

// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
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
