package main

import "fmt"

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
*/

func main() {
	// arr := []int{3, 0, 1}
	arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	n := len(arr)
	sum := n * (n + 1) / 2
	c_sum := 0
	for i := 0; i < n; i++ {
		c_sum += arr[i]
	}
	fmt.Println(sum - c_sum)
}
