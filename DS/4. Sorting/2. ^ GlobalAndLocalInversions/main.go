package main

import "fmt"

/*
775 -> https://leetcode.com/problems/global-and-local-inversions/description/
You are given an integer array nums of length n which represents a permutation of all the integers in the range [0, n - 1].
The number of global inversions is the number of the different pairs (i, j) where:
0 <= i < j < n
nums[i] > nums[j]
The number of local inversions is the number of indices i where:
0 <= i < n - 1
nums[i] > nums[i + 1]
Return true if the number of global inversions is equal to the number of local inversions.
Example 1:
Input: nums = [1,0,2]
Output: true
Explanation: There is 1 global inversion and 1 local inversion.
Example 2:
Input: nums = [1,2,0]
Output: false
Explanation: There are 2 global inversions and 1 local inversion.

O(nlog n) sol^n
	Global Inversion -> MergeSort -> O(nlogn)
	LocalInvesrsion -> O(n)

O(n) sol^n
	All local inversions are global inversions.
	If the number of global inversions is equal to the number of local inversions,
	it means that all global inversions in permutations are local inversions.
	It also means that we can not find A[i] > A[j] with i+2<=j.
	In other words, max(A[i]) < A[i+2]
*/

func main() {
	nums := []int{0, 2, 3, 1}
	fmt.Println(isIdealPermutationOptimal(nums))
	fmt.Println(isIdealPermutation(nums))

}
func isIdealPermutationOptimal(nums []int) bool {
	cmax := 0
	for i := 0; i < len(nums)-2; i++ {
		if cmax < nums[i] {
			cmax = nums[i]
		}
		if cmax > nums[i+2] {
			return false
		}
	}
	return true
}
func isIdealPermutation(nums []int) bool {
	localInversionCount := 0
	// globalInversionCount := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			localInversionCount++
		}
	}
	globalInversionCount := mergeSort(nums, 0, len(nums)-1)

	// fmt.Println(globalInversionCount)
	// fmt.Println(localInversionCount)
	if globalInversionCount == localInversionCount {
		return true
	}
	return false
}

func mergeSort(nums []int, s, e int) int {
	if s >= e {
		return 0
	}
	m := (s + e) / 2
	c1 := mergeSort(nums, s, m)
	c2 := mergeSort(nums, m+1, e)
	c3 := merge(nums, s, e)
	// fmt.Println(nums, s, e, c1, c2, c3, c1+c2+c3)
	return c1 + c2 + c3
}

func merge(nums []int, s, e int) int {
	m := (s + e) / 2
	i, j := s, m+1
	temp := []int{}
	cnt := 0
	for i <= m && j <= e {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			cnt += m - i + 1
			j++
		}
	}
	for i <= m {
		temp = append(temp, nums[i])
		i++
	}
	for j <= e {
		temp = append(temp, nums[j])
		j++
	}
	// nums = temp
	copy(nums[s:e+1], temp)
	return cnt
}
