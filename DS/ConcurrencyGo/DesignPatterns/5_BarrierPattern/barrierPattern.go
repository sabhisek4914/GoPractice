package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Barrier pattern in GO is a concurrency pattern which is used to synchronize a group of go routines, ensuring that all
 goroutines reach a certain point before they can proceed any further.
This is particularly useful when multiple tasks must complete their work before moving to the next phase of execution.

It can be achived using sync.WaitGroup and sync.Cond
*/

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // Utilize all CPU cores
	fmt.Printf("Running on %d CPU cores\n", runtime.NumCPU())

	const numWorkers = 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	barrier := func(workerID int) {
		defer wg.Done()
		fmt.Printf("Worker %d is performing work.\n", workerID)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d has reached the barrier.\n", workerID)
	}

	for i := 1; i <= numWorkers; i++ {
		go barrier(i)
	}

	wg.Wait() // Wait for all workers to reach the barrier.
	fmt.Println("All workers have reached the barrier, Proceeding further")
}
