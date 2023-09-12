package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 5 // Replace with the number for which you want to calculate the factorial
	resultCh := make(chan uint64)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			resultCh <- uint64(num)
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	fmt.Println(resultCh)
	finalResult := uint64(1)
	for num := range resultCh {
		finalResult *= num
	}

	fmt.Printf("Factorial of %d: %d\n", n, finalResult)
}
