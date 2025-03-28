package main

import "fmt"

/*
162. Find Peak Element -> https://leetcode.com/problems/find-peak-element/description/
A peak element is an element that is strictly greater than its neighbors.
Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.
You must write an algorithm that runs in O(log n) time.

Example 1:
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:
Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.
*/

/*
why  nums[m] >= nums[m+1] ??
peak elemnt mean a[i-1]<a[i]>a[i+1]

now my binary search patter is in this manner
	if predicate(){
		e=m
	}else{
		s=m+1
	}

Now on TRUE my e pointer is reducing on FALSE s pointer is increasing
	-> So to reach peak on TRUE
		 nums[m] >= nums[m+1] -> bcz this says on right of m we have smaller elements that we don't need -> so we are making e=m -> right space is rejected
*/

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}

func findPeakElement(nums []int) int {
	s, e := 0, len(nums)-1
	for s < e {
		m := s + (e-s)/2
		if nums[m] >= nums[m+1] {
			e = m
		} else {
			s = m + 1
		}
	}
	return s
}
