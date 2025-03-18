package main

import "fmt"

/*
Find the kth Smallest Element
Input: nums = [3,1,5,6,4], k = 2
Output: 3

Approach1 -> MaxHeap/MinHeap -> PriorityQueue -> O(nlogk)
Approach2 -> Sort -> arr[k] -> O(nlogn)
Approach2 -> QuickSelect -> QuickSort on Half and BinarySearch -> O(n)
*/

func main() {
	arr := []int{3, 1, 5, 6, 4}
	k := 1
	op := quickSelect(arr, 0, len(arr)-1, k-1)
	fmt.Println(op)
}

func quickSelect(arr []int, s, e, k int) int {
	if s >= e {
		return arr[s]
	}
	p := generatePivot(arr, s, e)
	if p == k {
		return arr[p]
	} else if p > k {
		return quickSelect(arr, s, p-1, k)
	} else {
		return quickSelect(arr, p+1, e, k)
	}
	// return -1
}

func generatePivot(arr []int, s, e int) int {
	i := s - 1
	for j := s; j < e; j++ {
		if arr[j] < arr[e] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		// j++
	}
	arr[i+1], arr[e] = arr[e], arr[i+1]
	return i + 1
}
