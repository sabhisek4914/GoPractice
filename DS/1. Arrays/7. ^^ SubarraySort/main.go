package main

import (
	"fmt"
	"math"
)

// 581 -> https://leetcode.com/problems/shortest-unsorted-continuous-subarray/description/
// inp:={1,2,3,4,5,8,6,7,9,10,11}
//		 0,1,2,3,4,5,6,7,8,9,10
//op: [5,7] smallest subarray upon sorting sorts the entire array
func unordered(pos int, arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	if pos == 0 {
		return arr[pos] > arr[pos+1]
	}
	if pos == len(arr)-1 {
		return arr[pos] < arr[pos-1]
	}
	return arr[pos] > arr[pos+1] || arr[pos] < arr[pos-1]
}
func main() {
	input := []int{1, 2, 3, 4, 5, 8, 6, 7, 9, 10, 11}
	small, large := math.MaxInt64, math.MinInt64
	for i := 0; i < len(input); i++ {
		if unordered(i, input) {
			if input[i] > large {
				large = input[i]
			}
			if input[i] < small {
				small = input[i]
			}
		}
	}
	fmt.Println(small, large)
	s, e := 0, len(input)-1
	for s < len(input) && input[s] < small {
		s++
	}
	for e >= 0 && input[e] > large {
		e--
	}
	fmt.Println("POSITIONs:", s, e)
}
