package main

// Given an array of integers, find the length of the longest sub-sequence such that elements in the subsequence are consecutive integers,
// the consecutive numbers can be in any order.
// Examples:
// Input: arr[] = {1, 9, 3, 10, 4, 20, 2}
// Output: 4

// Explanation: The subsequence 1, 3, 4, 2 is the longest subsequence of consecutive elements
// Input: arr[] = {36, 41, 56, 35, 44, 33, 34, 92, 43, 32, 42}
// Output: 5

// Explanation: The subsequence 36, 35, 33, 34, 32 is the longest subsequence of consecutive elements.

// func fibonci(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	return fibonci(n-1) + fibonci(n-2)
// }
import "fmt"

func main() {
	arr := []int{1, 9, 3, 10, 4, 20, 2}
	mp := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = true
	}
	large := int32(-2147483648)
	for i := 0; i < len(arr); i++ {
		ctr := int32(0)
		j := arr[i]
		for mp[j] {
			ctr++
			j++
		}
		if ctr > large {
			large = ctr
		}
	}
	fmt.Println(arr)
	fmt.Println(large)

}
