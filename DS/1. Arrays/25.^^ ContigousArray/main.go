package main

import "fmt"

/*
525. Contiguous Array -> https://leetcode.com/problems/contiguous-array/description/
Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.
Example 1:
Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
Example 2:
Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

PrefixSum and Hashing

*/

// findMaxLength returns the maximum length of a contiguous subarray
// with an equal number of 0 and 1 using prefix sum.
func findMaxLength(nums []int) int {
	// Map to store the first occurrence of a prefix sum.
	// Initialize with sum 0 at index -1.
	prefixMap := map[int]int{0: -1}

	maxLength := 0
	sum := 0

	// Iterate through the array.
	for i, num := range nums {
		// Treat 0 as -1, and 1 as +1.
		if num == 0 {
			sum += -1
		} else {
			sum += 1
		}

		// If the same prefix sum has been seen before,
		// it means the subarray between the previous index and the current index is balanced.
		if prevIndex, exists := prefixMap[sum]; exists {
			length := i - prevIndex
			if length > maxLength {
				maxLength = length
			}
		} else {
			// Store the first occurrence of this prefix sum.
			prefixMap[sum] = i
		}
	}

	return maxLength
}

func main() {
	// Example usage:
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	fmt.Println("Maximum length of contiguous subarray with equal number of 0 and 1:", findMaxLength(nums))
}
