package main

import "fmt"

/*
LeetCode 169. https://leetcode.com/problems/majority-element/description/
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element
always exists in the array.
Example 1:
Input: nums = [3,2,3]
Output: 3
Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2

check for Boyer-Moore Majority Voting Algorithm -> TO achive
timeComplexity: O(n)
spaceComplexity: O(1)

*/

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	op1 := majorityElementMapSolution(arr)
	/****IMP*****/
	op2 := majorityElementSpaceOptimialSolution(arr)
	fmt.Println(op1)
	fmt.Println(op2)
}

func majorityElementMapSolution(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	large := 0
	largeKey := 0
	for k, v := range mp {
		if v > large {
			large = v
			largeKey = k
		}
	}
	return largeKey
}

func majorityElementSpaceOptimialSolution(nums []int) int {
	majorityElement, ctr := nums[0], 1
	for i := 0; i < len(nums); i++ {
		if ctr == 0 {
			majorityElement, ctr = nums[i], 1
		} else {
			if nums[i] == majorityElement {
				ctr++
			} else {
				ctr--
			}
		}
	}
	return majorityElement
}
