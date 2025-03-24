package main

import (
	"fmt"
	"math"
)

/*
https://www.youtube.com/watch?v=VFC1oxkn5-E&list=PLICVjZ3X1AcYYdde4GTp79zfdp_VACSkX&index=7

33. Search in Rotated Sorted Array ->  https://leetcode.com/problems/search-in-rotated-sorted-array/description/
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

See copy to understand better(Pattern based also)
*/

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	// x:=search(nums, target)
	fmt.Println(search(nums, target))
	fmt.Println(searchPatternBased(nums, target))
}

func searchPatternBased(nums []int, target int) bool {
	s, e := 0, len(nums)-1
	for s < e {
		m := s + (e-s)/2
		if predicate(nums, m, target) {
			e = m
		} else {
			s = m + 1
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

func search(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		m := (s + e) / 2
		if nums[m] == target {
			return m
		}
		if nums[s] <= nums[m] {
			if target >= nums[s] && target <= nums[m] {
				e = m - 1
			} else {
				s = m + 1
			}
		} else {
			if target <= nums[e] && target >= nums[m] {
				s = m + 1
			} else {
				e = m - 1
			}
		}
	}
	return -1
}
