package main

import "fmt"

// Input: arr = [2,1,4,7,3,2,5]
// Output: 5
// Explanation: The largest mountain is [1,4,7,3,2] which has length 5.

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}
	fmt.Println(arr)
	large := 0
	if len(arr) < 3 {
		fmt.Println(0)
	} else {
		for i := 1; i < len(arr)-1; {
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				j := i
				ctr := 1
				for j > 0 && arr[j] > arr[j-1] {
					ctr++
					j--
				}
				for i < len(arr)-1 && arr[i] > arr[i+1] {
					ctr++
					i++
				}
				if large < ctr {
					large = ctr
				}
			} else {
				i++
			}
		}
		fmt.Println(large)

	}
}
