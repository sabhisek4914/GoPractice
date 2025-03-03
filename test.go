package main

import "fmt"

func modifySlice(a []int) {
	s := a
	s[0] = 7
	s = append(s, 100, 101, 102) // Creates a new slice, does not affect original
	s[0] = 9
	s[3] = 8
	fmt.Println(s)
}

func main() {
	nums := make([]int, 8, 10)
	// nums = []int{1, 2, 3}
	nums[0] = 1
	nums[1] = 2
	nums[3] = 3
	modifySlice(nums)
	fmt.Println(nums) // Output: [1 2 3] (No change!)
}
