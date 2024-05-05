package main

import "fmt"

// inp:{1,2,3,4,5}
// op:{120,60,40,30,24}

//Brute Force
// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	left := make([]int, len(arr))
// 	right := make([]int, len(arr))
// 	op := make([]int, len(arr))
// 	left[0] = 1
// 	right[len(right)-1] = 1
// 	temp := 1
// 	for i := 1; i < len(arr); i++ {
// 		left[i] = temp
// 		temp *= arr[i]
// 	}
// 	temp = 1
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		right[i] = temp
// 		temp *= arr[i]
// 	}
// 	fmt.Println(left)
// 	fmt.Println(right)
// 	for j := 0; j < len(arr); j++ {
// 		op[j] = left[j] * right[j]
// 	}
// 	fmt.Println(op)
// }

func main() {
	arr := []int{1, 2, 3, 4, 5}
	op := make([]int, len(arr))
	temp := 1
	for i := 0; i < len(arr); i++ {
		op[i] = temp
		temp *= arr[i]
	}
	temp = 1

	for i := len(arr) - 1; i >= 0; i-- {
		op[i] *= temp
		temp *= arr[i]
	}
	fmt.Println(op)
}
