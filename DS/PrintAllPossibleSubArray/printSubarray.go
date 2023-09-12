package main

import "fmt"

func generateSubsequences(arr []int, index int, subsequence []int) {
	if index == len(arr) {
		fmt.Println(subsequence)
		return
	}

	// Exclude the current element
	generateSubsequences(arr, index+1, subsequence)

	// Include the current element
	subsequence = append(subsequence, arr[index])
	generateSubsequences(arr, index+1, subsequence)
}

func main() {
	arr := []int{1, 2, 3}
	//op: [[],[1],[1,2],[1,2],[1,2,3],[2],[2,3]]
	subarray := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr); j++ {
			a := arr[i:j]
			subarray = append(subarray, a)
		}
	}
	fmt.Println("SubArray Is: ", subarray)
	var subsequence []int
	generateSubsequences(arr, 0, subsequence)
}
