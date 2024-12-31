package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
The Sleeping Barber Problem is a classic synchronization problem that involves coordinating a barber, customers, and a waiting area.
The barber sleeps if no customers are present and wakes up when a customer arrives.
The challenge is to manage access to the barber's chair and waiting room in a way that avoids race conditions.
*/

const (
	waitingRoomSize = 3  // Max customers in waiting room
	numCustomers    = 20 // Total number of customers
)

var wg sync.WaitGroup

type BarberShop struct {
	chair       chan struct{} // Barber chair
	waitingRoom chan int      // Waiting room seats
	barberDone  chan struct{} // Signal for barber completing a haircut
	// customerReady chan int      // Signal for a customer being ready
	closeSaloon chan struct{} // Signal for barber to close the shop
}

// Barber Routine
func (shop *BarberShop) Barber() {
	for {
		select {
		case customerID := <-shop.waitingRoom: // Customer wakes barber
			fmt.Printf("Barber invites Customer %d to the chair.\n", customerID)
			shop.chair <- struct{}{} // Barber allows customer to sit
			fmt.Printf("Barber is cutting Customer %d's hair...\n", customerID)
			time.Sleep(time.Second)
			fmt.Printf("Barber has finished cutting Customer %d's hair.\n", customerID)
			shop.barberDone <- struct{}{}
		case <-shop.closeSaloon: // Close salon as no more customers are coming
			fmt.Println("Closing the salon as no more customers are coming.")
			return
		default:
			fmt.Println("Barber is sleeping...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// Customer Routine
func (shop *BarberShop) Customer(id int) {
	defer wg.Done()
	// time.Sleep(time.Millisecond * 500)
	fmt.Printf("Customer %d arrived.\n", id)

	select {
	case shop.waitingRoom <- id:
		fmt.Printf("Customer %d is waiting.\n", id)
		// shop.customerReady <- id // Notify barber
		<-shop.chair // Occupy barber's chair
		fmt.Printf("Customer %d is getting a haircut.\n", id)
		<-shop.barberDone // Wait for haircut to finish
		fmt.Printf("Customer %d is leaving.\n", id)
	default:
		fmt.Printf("Customer %d left as waiting room is full.\n", id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	shop := BarberShop{
		chair:       make(chan struct{}, 1),
		waitingRoom: make(chan int, waitingRoomSize), // Store customer IDs in the waiting room
		barberDone:  make(chan struct{}),
		// customerReady: make(chan int), // Signal barber with customer ID
		closeSaloon: make(chan struct{}),
	}

	go shop.Barber()

	for i := 1; i <= numCustomers; i++ {
		wg.Add(1)
		go shop.Customer(i)
		time.Sleep(time.Millisecond * 200) // Simulate arrival of customers
	}

	wg.Wait()

	// Signal barber to close the shop
	shop.closeSaloon <- struct{}{}
	// Wait for barber to process the close signal
	time.Sleep(time.Second)
	fmt.Println("Barber shop is closed.")
}
