package main

import (
	"fmt"
	"sync"
)

/* The worker pool pattern involves creating a group of worker goroutines to process tasks concurrently,
limiting the number of simultaneous operations. This pattern is valuable when you have a large number of
tasks to execute.
Applications:
1. Handling incoming HTTP requests in a web server.
2. Processing images concurrently.
*/

func main() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	//Start worker routines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	//Enqueue Jobs
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	//Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	//Collect the results
	for results := range results {
		fmt.Println("Results: ", results)
	}
}

func worker(workerID int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %v Procesing Job: %v\n", workerID, job)
		results <- job * 2
	}
}
