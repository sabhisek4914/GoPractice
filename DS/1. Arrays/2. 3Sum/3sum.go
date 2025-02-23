package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/description/
//{12, 3, 4, 1, 6, 9}, sum = 24;
// O(n^2) time
// O(1) space

func processThreeSum(arr []int, sum int) {
	sort.Ints(arr)
	fmt.Println("arr: ", arr)
	for i := 0; i < len(arr)-2; i++ {
		remain := sum - arr[i]
		s := i + 1
		e := len(arr) - 1
		// fmt.Printf("SA %d", s)
		for s < e {
			if arr[s]+arr[e] == remain {
				fmt.Println(arr[i], arr[s], arr[e])
				s++
				e--
			} else if arr[s]+arr[e] < remain {
				s += 1
			} else {
				e -= 1
			}
		}
	}
}

func processThreeSumMap(arr []int, sum int) {
	for i := 0; i < len(arr); i++ {
		remain := sum - arr[i]
		mp := make(map[int]int)
		for j := 0; j < len(arr); j++ {
			r2 := remain - arr[j]
			if r2 != arr[i] {
				if val, ok := mp[r2]; ok {
					// if val == 1 {
					fmt.Println(arr[i], arr[j], r2, ok, val)
					// } else if val > 1 {
					// 	val++
					// 	mp[r2] = val
					// }
				} else {
					mp[arr[j]] = 0
				}
			}
		}
	}
}

func main() {
	arr := []int{12, 3, 4, 1, 6, 9}
	sum := 24
	processThreeSum(arr, sum)
	processThreeSumMap(arr, sum)
}
