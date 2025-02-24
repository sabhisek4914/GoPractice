package main

import (
	"container/heap"
	"fmt"
)

// MinHeap type (implements heap.Interface)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap: Smallest element on top
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push inserts an element into the heap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the smallest element
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[0]   // Smallest element
	*h = old[1:n] // Remove the root
	return x
}

// Function to find the K-th smallest element
func findKthSmallest(nums []int, k int) int {
	h := MinHeap(nums) // Convert slice to MinHeap
	heap.Init(&h)      // Heapify the slice (O(N))

	// Extract min k times
	var kthSmallest int
	for i := 0; i < k; i++ {
		kthSmallest = heap.Pop(&h).(int) // Pop the smallest element
	}

	return kthSmallest
}

func main() {
	nums := []int{7, 10, 4, 3, 20, 15}
	k := 3
	fmt.Println("K-th Smallest Element:", findKthSmallest(nums, k)) // Output: 7
}
