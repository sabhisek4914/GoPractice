package main

import "fmt"

// Write a function name findMaxConsecutiveOnes that could nums as argument. Given a binary array nums, return the maximum number of consecutive 1's in the array.
// Example 1: Input: nums = [1,1,0,1,1,1] Output: 3 Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
// Example 2: Input: nums = [1,0,1,1,0,1] Output: 2

func main() {
	arr := []int{1, 1, 0, 1, 1, 1}
	large := 0
	ctr := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			ctr++
		} else if arr[i] == 0 {
			if ctr > large {
				large = ctr
			}
			ctr = 0
		}
	}
	if ctr > large {
		large = ctr
	}
	fmt.Println(large)
}
