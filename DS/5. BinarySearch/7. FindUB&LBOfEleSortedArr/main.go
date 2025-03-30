package main

import (
	"fmt"
	"sort"
)

/*
34. Find First and Last Position of Element in Sorted Array -> https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.
Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

use sort.Search() prebuilt method to do the binarysearch in sorted[]
*/

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
	fmt.Println(searchRangeBuiltIn(nums, target))
}

//IMP
func searchRangeBuiltIn(nums []int, target int) []int {
	start := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := sort.Search(len(nums)-start, func(i int) bool {
		return nums[start+i] > target
	}) + start - 1
	return []int{start, end}
}

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	lb := lowerBoundSearch(nums, target)
	ub := upperBoundSearch(nums, target)
	return []int{lb, ub}
}

func lowerBoundSearch(nums []int, target int) int {
	s, e := 0, len(nums)-1
	if nums[s] == target {
		return s
	}
	for s < e {
		m := s + (e-s)/2
		if nums[m] >= target {
			e = m
		} else {
			s = m + 1
		}
	}
	if nums[s] == target {
		return s
	}
	return -1
}

func upperBoundSearch(nums []int, target int) int {
	s, e := 0, len(nums)-1
	if nums[e] == target {
		return e
	}
	for s < e {
		m := s + (e-s)/2
		if nums[m] > target {
			e = m
		} else {
			s = m + 1
			fmt.Println(s, e)
		}
	}
	if s != 0 && nums[s-1] == target {
		return s - 1
	}
	return -1
}
