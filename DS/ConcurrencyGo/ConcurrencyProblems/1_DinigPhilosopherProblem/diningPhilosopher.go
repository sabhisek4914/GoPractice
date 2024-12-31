package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent
philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they
 have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only
if it is not being used by another philosopher. After an individual philosopher finishes eating,
they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right
or the one on their left as they become available, but cannot start eating before getting both forks.
Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed.
The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve; i.e.,
each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may
want to eat or think.
*/

const numPhilosophersCount = 5

var wg sync.WaitGroup
var forks [numPhilosophersCount]sync.Mutex

func philosopher(id int, left, right *sync.Mutex) {
	defer wg.Done()
	//Assuming each philosopher stomach/hunnger is filled eating 3 times.
	for i := 0; i < 3; i++ {
		fmt.Printf("Philosopher %d is thinking\n", id)
		time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Printf("Philoshoper %d is hungry\n", id)
		left.Lock()
		right.Lock()
		fmt.Printf("Philoshoper %d is eating\n", id)
		time.Sleep(time.Millisecond * time.Duration(100))
		left.Unlock()
		right.Unlock()
		fmt.Printf("Philoshoper %d is done eating\n", id)
	}
	fmt.Printf("Philoshoper %d is full\n", id)
}
func main() {
	wg.Add(numPhilosophersCount)
	for i := 0; i < numPhilosophersCount; i++ {
		left := &forks[i]
		right := &forks[(i+1)%numPhilosophersCount]
		go philosopher(i, left, right)
	}
	wg.Wait()
}
