package main

import "fmt"

/*
560. Subarray Sum Equals K ->https://leetcode.com/problems/subarray-sum-equals-k/description/
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:
Input: nums = [1,1,1], k = 2
Output: 2
Example 2:
Input: nums = [1,2,3], k = 3
Output: 2

subarrays -> [1,2,3]
1
2
3
1,2
2,3
1,2,3
*/
func main() {
	nums := []int{1, 2, 3}
	k := 3
	x := subarraySum(nums, k)
	fmt.Println(x)
	x = subarraySumBruteForce(nums, k)
	fmt.Println(x)
}

func subarraySum(nums []int, k int) int {
	count := 0
	m := map[int]int{0: 1} // if the sub array first item == k
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			// in the map we have count of subsets which add to sum-k
			count += v
		}
		m[sum]++
	}
	return count
}

func subarraySumBruteForce(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if nums[i] == k {
			count++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				fmt.Println(i, j)
				count++
			}
		}
	}
	return count
}
