package main

import "math"

/*
1283. Find the Smallest Divisor Given a Threshold -> https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/description/
Given an array of integers nums and an integer threshold, we will choose a positive integer divisor, divide all the array by it, and sum the division's result. Find the smallest divisor such that the result mentioned above is less than or equal to threshold.
Each result of the division is rounded to the nearest integer greater than or equal to that element. (For example: 7/3 = 3 and 10/2 = 5).
The test cases are generated so that there will be an answer.

Example 1:
Input: nums = [1,2,5,9], threshold = 6
Output: 5
Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1.
If the divisor is 4 we can get a sum of 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2).
Example 2:
Input: nums = [44,22,33,11,1], threshold = 5
Output: 44

It is a BS problem?? bcz it is in a monotonous range(pattern)
what is the range of s & e
e := max(arr)
s:=1
*/

func main() {
	nums := []int{44, 22, 33, 11, 1}
	threshold := 5
	smallestDivisor(nums, threshold)
}
func smallestDivisor(nums []int, threshold int) int {
	s, e := 1, slices.Max(nums)
	for s < e {
		m := s + (e-s)/2
		if predicate(nums, threshold, m, s, e) {
			e = m
		} else {
			s = m + 1
		}
	}
	return s
}
func predicate(nums []int, threshold, m, s, e int) bool {
	if generateSum(nums, m) <= threshold {
		if m <= e {
			return true
		}
	}
	return false
}
func generateSum(nums []int, m int) int {
	sum := 0
	for _, v := range nums {
		sum += int(math.Ceil(float64(v) / float64(m)))
	}
	return sum
}
