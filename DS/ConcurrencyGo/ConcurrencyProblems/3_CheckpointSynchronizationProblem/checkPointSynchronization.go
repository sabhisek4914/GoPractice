package main

import (
	"fmt"
	"time"
)

/*
Checkpoint Synchronization in Go is a concurrency pattern where multiple goroutines coordinate their progress to a defined
 point (checkpoint) before proceeding. It ensures that all goroutines reach a particular state or complete a task before
 continuing. This is often achieved using synchronization mechanisms like WaitGroups, Channels, or Barriers.
*/

func main() {
	numWorkers := 3
	ch := make(chan struct{}, numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	//If range ch then i had to close() channel first
	for i := 0; i < numWorkers; i++ {
		<-ch
	}
	fmt.Println("All workers reached the checkpoint.")

}

func worker(id int, ch chan struct{}) {
	fmt.Printf("Worker:%d started\n", id)
	time.Sleep(time.Second * time.Duration(id))
	fmt.Printf("Worker:%d Reached Checkpoint\n", id)
	ch <- struct{}{} //Signal Completion
}

/*
Using sync package
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Mark this goroutine as done
    fmt.Printf("Worker %d: Started\n", id)
    time.Sleep(time.Second * time.Duration(id)) // Simulate work
    fmt.Printf("Worker %d: Reached checkpoint\n", id)
}

func main() {
    var wg sync.WaitGroup
    numWorkers := 3

    for i := 1; i <= numWorkers; i++ {
        wg.Add(1) // Add a goroutine to the WaitGroup
        go worker(i, &wg)
    }

    wg.Wait() // Wait for all goroutines to reach the checkpoint
    fmt.Println("All workers reached the checkpoint. Proceeding...")
}


Using Barrier:
package main

import (
    "fmt"
    "time"
)

func worker(id int, ch chan struct{}) {
    fmt.Printf("Worker %d: Started\n", id)
    time.Sleep(time.Second * time.Duration(id)) // Simulate work
    fmt.Printf("Worker %d: Reached checkpoint\n", id)
    ch <- struct{}{} // Signal completion
}

func main() {
    numWorkers := 3
    ch := make(chan struct{}, numWorkers) // Buffered channel

    for i := 1; i <= numWorkers; i++ {
        go worker(i, ch)
    }

    for i := 0; i < numWorkers; i++ {
        <-ch // Wait for each worker to signal completion
    }

    fmt.Println("All workers reached the checkpoint. Proceeding...")
}

*/
