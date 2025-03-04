package main

import "fmt"

/*
26. Remove duplicates from sorted array using same array -> https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 3, 5}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
