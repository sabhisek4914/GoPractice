package main

import (
	"fmt"
	"sync"
)

/*
The fan-out/fan-in pattern involves distributing tasks to multiple worker goroutines (fan-out)
and then aggregating their results (fan-in).
It’s useful for parallelizing tasks and combining their outcomes.
Application:
1. Web scraping multiple websites concurrently and merging the results.
2. Aggregating data from multiple sensors in IoT applications.
*/

func main() {
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))

	//Fan-out launc multiple worker goroutines
	numWorkers := 3
	results := make(chan int, len(data))

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for num := range input {
				fmt.Printf("i:%d num:%d \n", n, num)
				results <- num * 2
			}
		}(i)
		// time.Sleep(2 * time.Second)
	}

	for _, d := range data {
		input <- d
	}
	close(input)

	//Fan-in: Aggregate results from workers
	go func() {
		wg.Wait()
		close(results)
	}()

	//Process aggregated results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
