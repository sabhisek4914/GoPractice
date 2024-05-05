package main

import "fmt"

/*
LEET CODE: 283
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
*/

func moveZeroes(arr []int) {
	l := 0
	r := 0
	for r < len(arr) {
		if arr[r] != 0 {
			//swap
			arr[l], arr[r] = arr[r], arr[l]
			l++
		}
		r++
	}
	fmt.Println(arr)
}
func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}
