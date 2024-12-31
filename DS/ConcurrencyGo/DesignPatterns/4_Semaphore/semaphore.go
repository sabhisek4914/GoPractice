package main

import (
	"fmt"
	"sync"
	"time"
)

/*The semaphore pattern in go is a concurrency control mechanism that limits the number of go routines that
 access a shared resource simultenously.
How Semaphore Works
Tokens: A semaphore uses a counter (or tokens) to track the number of available "permits."
		Each goroutine acquires a token before entering the critical section and releases it after completing its work.
Blocking: If no tokens are available, the semaphore blocks the goroutine until a token is released.
*/

func main() {
	const maxConcurrent = 3 // Limit to 3 routines
	semaphore := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	worker := func(id int) {
		defer wg.Done()
		semaphore <- struct{}{} //accquire a token
		fmt.Printf("Worker %d started\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finished\n", id)
		<-semaphore //Release token
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}

/*semaphore eusing sync.cond
package main

import (
    "fmt"
    "sync"
    "time"
)

type Semaphore struct {
    permits int
    cond    *sync.Cond
}

func NewSemaphore(maxPermits int) *Semaphore {
    return &Semaphore{
        permits: maxPermits,
        cond:    sync.NewCond(&sync.Mutex{}),
    }
}

func (s *Semaphore) Acquire() {
    s.cond.L.Lock()
    for s.permits == 0 {
        s.cond.Wait() // Wait until a permit is available
    }
    s.permits--
    s.cond.L.Unlock()
}

func (s *Semaphore) Release() {
    s.cond.L.Lock()
    s.permits++
    s.cond.Signal() // Notify waiting goroutines
    s.cond.L.Unlock()
}

func main() {
    semaphore := NewSemaphore(3) // Max 3 concurrent goroutines

    var wg sync.WaitGroup

    worker := func(id int) {
        defer wg.Done()
        semaphore.Acquire()
        fmt.Printf("Worker %d started\n", id)
        time.Sleep(2 * time.Second) // Simulate work
        fmt.Printf("Worker %d finished\n", id)
        semaphore.Release()
    }

    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go worker(i)
    }

    wg.Wait()
}
Comparison: Semaphore vs Mutex
Aspect		Semaphore									Mutex
Concurrency	Allows multiple goroutines (up to a limit).	Allows only one goroutine.
Use Case	Resource pooling, rate limiting.			Critical section protection.
Blocking	Blocks when no tokens are available.		Blocks until lock is released.
Granularity	More granular control (permits > 1).		Binary control (lock/unlock).
*/
