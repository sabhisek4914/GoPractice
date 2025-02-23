package main

/*
31. Next Permutation-> https://leetcode.com/problems/next-permutation/description/
[1,2,3] -> All permutations are [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1]
Lexographical order is
[1,2,3]
[1,3,2]
[2,1,3]
[2,3,1]
[3,1,2]
[3,2,1]

Lexgraphical Next PErmutation is ????
Example 1:
Input: nums = [1,2,3]
Output: [1,3,2]
Example 2:
Input: nums = [3,2,1]
Output: [1,2,3]
Example 3:
Input: nums = [1,1,5]
Output: [1,5,1]

Approach:
1. Find the 1st decreasing element from Right
2. If found in 1. find the next larger element
3. Sort from 1. to end
*/

import (
	"fmt"
	"sort"
)

// nextPermutation finds the next lexicographical permutation
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Step 1: Find the first decreasing element from the right -> if [1,2,3] then 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	//if i==-1 -> [] is in decreasing order i.e. [3,2,1]
	if i >= 0 { // If we found such an index
		// Step 2: Find the next larger element to swap with
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// Swap nums[i] and nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 3: Reverse the suffix (nums[i+1:])
	sort.Ints(nums[i+1:])
}

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [2, 1, 3]
}
