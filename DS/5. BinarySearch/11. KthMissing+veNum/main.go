package main

import "fmt"

/*
1539. Kth Missing Positive Number -> https://leetcode.com/problems/kth-missing-positive-number/description/
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.
Return the kth positive integer that is missing from this array.

Example 1:
Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
Example 2:
Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
*/
func main() {
	nums := []int{44, 22, 33, 11, 1}
	threshold := 5
	findKthPositive(nums, threshold)
}

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left)
	return left + k
}
