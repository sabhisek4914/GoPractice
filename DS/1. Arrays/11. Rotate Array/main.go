package main

import "fmt"

/*
LEET CODE: 189
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
*/
func rotateRightArr(nums []int, k int) {
	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	fmt.Println(nums)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotateRightArr(nums, k)
	fmt.Println(nums)

}
