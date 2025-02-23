package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 3, 4, 5}
	var nums [][]int
	nums = append(nums, nums1, nums2)
	fmt.Println(nums)
	fmt.Println(len(nums), len(nums[0]))
}
