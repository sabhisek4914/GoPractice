package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	paper = iota
	grass
	match
)

var smokeMap = map[int]string{
	paper: "paper",
	grass: "grass",
	match: "match",
}

var names = map[int]string{
	paper: "Sandy",
	grass: "Apple",
	match: "Daisy",
}

type Table struct {
	paper chan int
	grass chan int
	match chan int
}

func arbitrate(t *Table, smokers [3]chan int) {
	for {
		time.Sleep(time.Millisecond * 500)
		next := rand.Intn(3)
		fmt.Printf("Table chooses %s: %s\n", smokeMap[next], names[next])
		switch next {
		case paper:
			t.grass <- 1
			t.match <- 1
		case grass:
			t.paper <- 1
			t.match <- 1
		case match:
			t.grass <- 1
			t.paper <- 1
		}
		for _, smoker := range smokers {
			smoker <- next
		}
		wg.Add(1)
		wg.Wait()
	}
}

func smoker(t *Table, name string, smokes int, signal chan int) {
	var chosen = -1
	for {
		chosen = <-signal // blocks

		if smokes != chosen {
			continue
		}

		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		select {
		case <-t.paper:
		case <-t.grass:
		case <-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		time.Sleep(10 * time.Millisecond)
		select {
		case <-t.paper:
		case <-t.grass:
		case <-t.match:
		}
		fmt.Printf("Table: %d grass: %d match: %d\n", len(t.paper), len(t.grass), len(t.match))
		fmt.Printf("%s smokes a cigarette\n", name)
		time.Sleep(time.Millisecond * 500)
		wg.Done()
		time.Sleep(time.Millisecond * 100)
	}
}

const LIMIT = 1

var wg *sync.WaitGroup

func main() {
	wg = new(sync.WaitGroup)
	table := new(Table)
	table.match = make(chan int, LIMIT)
	table.paper = make(chan int, LIMIT)
	table.grass = make(chan int, LIMIT)
	var signals [3]chan int
	// three smokers
	for i := 0; i < 3; i++ {
		signal := make(chan int, 1)
		signals[i] = signal
		go smoker(table, names[i], i, signal)
	}
	fmt.Printf("%s, %s, %s, sit with \n%s, %s, %s\n\n", names[0], names[1], names[2], smokeMap[0], smokeMap[1], smokeMap[2])
	arbitrate(table, signals)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// /*
// The Cigarette Smoker Problem is a classic synchronization problem in computer science. It involves three smokers and one agent
// who coordinates them. Each smoker requires three ingredients (tobacco, paper, and matches) to make a cigarette but has an infinite supply of only one.
// The agent places two random ingredients on the table, and the smoker who has the third ingredient makes a cigarette
// and smokes it.

// Smokers: Each smoker has an infinite supply of one ingredient:
// 	Smoker A: Has tobacco.
// 	Smoker B: Has paper.
// 	Smoker C: Has matches.
// Agent: Places two random ingredients on the table.
// Synchronization Challenge:
// 	Only the smoker with the missing ingredient can proceed.
// 	Other smokers must wait.
// */
// var wg sync.WaitGroup
// var tobaccoChan = make(chan struct{})
// var paperChan = make(chan struct{})
// var matchChan = make(chan struct{})

// var readyToSmoke = make(chan struct{})

// func agent() {
// 	time.Sleep(time.Millisecond * 500) // Simulate delay
// 	wg.Done()
// 	// Randomly choose two ingredients to place on the table
// 	for i := 0; i < 10; i++ { // Allow the agent to place ingredients 10 times
// 		switch rand.Intn(3) {
// 		case 0: //Place paper and matches
// 			fmt.Println("Agent places paper and matches on the table.")
// 			paperChan <- struct{}{}
// 			matchChan <- struct{}{}
// 		case 1: //Place paper and tobacco
// 			fmt.Println("Agent places paper and tobacco on the table.")
// 			paperChan <- struct{}{}
// 			tobaccoChan <- struct{}{}
// 		case 2: //Place tobacco and matches
// 			fmt.Println("Agent places tobacco and matches on the table.")
// 			tobaccoChan <- struct{}{}
// 			matchChan <- struct{}{}
// 		}

// 		//Wait for smoker to finish
// 		<-readyToSmoke
// 	}
// }

// func smoker(name string, hasIngredient chan struct{}, ingredient1 chan struct{}, ingredient2 chan struct{}) {
// 	defer wg.Done()
// 	select {
// 	case <-hasIngredient:
// 		<-ingredient1
// 		<-ingredient2
// 		fmt.Printf("%s makes and smokes a cigarette.\n", name)
// 		time.Sleep(time.Millisecond * 1000) // Simulate smoking time
// 		readyToSmoke <- struct{}{}          // Notify the agent
// 	default:
// 		// Avoid blocking indefinitely
// 		return
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go agent()

// 	//Smoker
// 	wg.Add(3)
// 	go smoker("Smoker A (Tobacco)", tobaccoChan, paperChan, matchChan)
// 	go smoker("Smoker B (Paper)", paperChan, tobaccoChan, matchChan)
// 	go smoker("Smoker C (Matches)", matchChan, tobaccoChan, paperChan)

// 	wg.Wait()
// 	fmt.Println("All smoking rounds completed.")
// }
