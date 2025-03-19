package main

import "fmt"

/*
Count the freq of an element in sorted []
inp: [0,1,1,1,2,2,2,3,4,4,5,10]
key: 2

op: 3 (2 2's are present)
linear search -> O(n)
binary search -> O(log n)

Approach: Find lowerBound of 2 and upper bound of 2 their difference is the count
*/

func main() {
	arr := []int{0, 1, 1, 1, 2, 2, 2, 3, 4, 4, 5, 10}
	key := 2
	lb := searchLB(arr, key)
	ub := searchUB(arr, key)
	fmt.Println(ub, lb, ub-lb+1)
}

func searchLB(arr []int, k int) int {
	ans := -1
	s, e := 0, len(arr)-1
	for s <= e {
		m := (s + e) / 2
		if arr[m] == k {
			ans = m
			e = m - 1
		} else if arr[m] > k {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return ans
}

func searchUB(arr []int, k int) int {
	ans := -1
	s, e := 0, len(arr)-1
	for s <= e {
		m := (s + e) / 2
		if arr[m] == k {
			ans = m
			s = m + 1
		} else if arr[m] > k {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return ans
}
