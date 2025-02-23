package main

import "fmt"

/*
LeetCode 75
https://leetcode.com/problems/sort-colors/description/
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
You must solve this problem without using the library's sort function.

Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]

Thought Process:
	try to separate all 1s to left and 2s to right
	use swap logic
	while iterating towards end, 1s will be handled automatically nd [] is sorted.
	to handle left nd right use 2 vars for conifrming 0 & 2, one more var to iterate overall.


TimeComplexity: O(n)
SpaceComplexity: O(n)
*/
func main() {
	fmt.Println("nums")
	// nums := []int{2, 1, 0, 0, 1, 2}
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		for i >= l && i <= r && (nums[i] == 0 || nums[i] == 2) {
			if nums[i] == 0 {
				// swap i,l
				nums[i], nums[l] = nums[l], nums[i]
				l++ //l postion is secured
			}
			if nums[i] == 2 {
				// swap i,r
				nums[i], nums[r] = nums[r], nums[i]
				r-- //r postion is secured
			}
		}
		// fmt.Println(nums)
	}
}
