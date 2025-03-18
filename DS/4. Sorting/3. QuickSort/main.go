package main

import "fmt"

/*
Quick Sort
inp: [10,5,2,0,7,6,4]
avg Case -> O(nlogn)
worstCase -> O(n^2)
*/

func main() {
	arr := []int{10, 5, 2, 0, 7, 6, 4}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, s, e int) {
	if s >= e {
		return
	}
	p := processPivot(arr, s, e)
	quickSort(arr, s, p-1)
	quickSort(arr, p+1, e)
}
func processPivot(arr []int, s, e int) int {
	pivot := arr[e]
	i, j := s-1, s
	for j <= e-1 {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}
	arr[i+1], arr[e] = pivot, arr[i+1]
	return i + 1
}
