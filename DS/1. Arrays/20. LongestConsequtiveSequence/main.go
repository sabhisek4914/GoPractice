package main

import "fmt"

/*
128. Longest Consecutive Sequence -> https://leetcode.com/problems/longest-consecutive-sequence/description/
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Approach1: sort -> use map to check element exists or not -> O(nlogn)
Approach2: use map -> iterate on the map to get an element whose -1 of it doesn't exist -> then take count andincrement and check for max O(n)
Approach3: use map -> use map, iterate over array and do the same. O(n)
*/

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num+1
		for set[temp] {
			sequence++
			temp++
		}

		if res < sequence {
			res = sequence
		}
		if sequence > len(nums)/2 {
			break
		}
	}
	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	cnt := longestConsecutive(nums)
	fmt.Println(cnt)
}
