package main

import (
	"fmt"
	"math"
)

/*
https://www.youtube.com/watch?v=VFC1oxkn5-E&list=PLICVjZ3X1AcYYdde4GTp79zfdp_VACSkX&index=7
81. Search in Rotated Sorted Array II -> https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
You must decrease the overall operation steps as much as possible.

Example 1:
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
*/

func main() {
	// nums := []int{2, 5, 6, 0, 0, 1, 2}
	nums := []int{1, 1, 1, 1, 1, 2, 1, 1}
	target := 2
	// x:=search(nums, target)
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) bool {
	s, e := 0, len(nums)-1
	for s < e {
		for s < e && nums[s] == nums[s+1] {
			s++
		}
		for s < e && nums[e] == nums[e-1] {
			e--
		}
		m := s + (e-s)/2
		if predicate(nums, m, target) {
			e = m
		} else {
			if s < len(nums)-1 {
				s = m + 1
			}
		}
	}
	fmt.Println(s, e)
	if nums[s] == target {
		return true
	}
	return false
}

func predicate(nums []int, m, target int) bool {
	val := 0
	if (target < nums[0]) == (nums[m] < nums[0]) {
		val = nums[m]
	} else {
		if target < nums[0] {
			val = math.MinInt32
		} else {
			val = math.MaxInt32
		}
	}
	return val >= target
}
