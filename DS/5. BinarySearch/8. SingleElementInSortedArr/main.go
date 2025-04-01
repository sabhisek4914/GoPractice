package main

import "fmt"

/*
540. Single Element in a Sorted Array -> https://leetcode.com/problems/single-element-in-a-sorted-array/description/
You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
Return the single element that appears only once.
Your solution must run in O(log n) time and O(1) space.

Example 1:
Input: nums = [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:
Input: nums = [3,3,7,7,10,11,11]
Output: 10
*/

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) int {
	s, e := 0, len(nums)-1
	for s < e {
		m := s + (e-s)/2
		if m%2 == 1 {
			m--
		}
		if nums[m] != nums[m+1] {
			e = m
		} else {
			s = m + 2
		}
		fmt.Println(s, e, m)
	}
	return nums[s]
}
