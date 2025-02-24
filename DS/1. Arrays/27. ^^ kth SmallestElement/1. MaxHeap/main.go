package main

import (
	"container/heap"
	"fmt"
)

/*
215. https://leetcode.com/problems/kth-largest-element-in-an-array/description/
Find the kth Smallest Element
Input: nums = [3,1,5,6,4], k = 2
Output: 3

Approch:
1. Sorting
2. Heap -> PriorityQueue
**https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc

MaxHeap
Building the heap: O(K)
Processing N elements: O((N-K) log K) (since we push & pop in O(log K))
Overall Complexity: O(N log K), which is efficient for large N.
*/

// MaxHeap type
type MaxHeap []int

//Interface Methods
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } //Max-Heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Push nd Pop for heap.Interface
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//Find kth smallest element
func findKthSmallest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0] //Top
}

func main() {
	nums := []int{7, 10, 4, 3, 20, 15}
	k := 3
	fmt.Println("K-th Smallest Element:", findKthSmallest(nums, k)) // Output: 7
}
