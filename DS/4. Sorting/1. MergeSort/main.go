package main

import "fmt"

/*
MergeSort ->Divide and conquer
O(nlogn)
Eg:
Input:	[10,5,2,0,7,6,4]
Output: [0,2,,4,5,6,7,10]
*/

func main() {
	arr := []int{0, 2, 3, 1}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, s, e int) {
	if s >= e {
		return
	}
	m := (s + e) / 2
	mergeSort(arr, s, m)
	mergeSort(arr, m+1, e)
	merge(arr, s, e)
	fmt.Println(arr)
}

func merge(arr []int, s, e int) {
	m := (s + e) / 2
	i := s
	j := m + 1
	temp := []int{}
	// k := 0
	for i <= m && j <= e {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++

		} else {
			temp = append(temp, arr[j])
			j++

		}
	}
	for j <= e {
		temp = append(temp, arr[j])
		j++
	}
	for i <= m {
		temp = append(temp, arr[i])
		i++
	}
	// arr = temp
	// Copy sorted elements back to original array
	copy(arr[s:e+1], temp)
}
