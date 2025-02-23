package main

import "fmt"

//42. https://leetcode.com/problems/trapping-rain-water/description/
// inp:=[0,1,0,2,1,0,1,3,2,1,2,1]
// l -> [0,1,1,2,2,2,2,3,3,3,3,3]
// r -> [3,3,3,3,3,3,3,3,2,2,2,1]
// op:= 6

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	left := make([]int, len(input))
	right := make([]int, len(input))
	left[0] = input[0]
	right[len(input)-1] = input[len(input)-1]
	for i := 1; i < len(input); i++ {
		if left[i-1] > input[i] {
			left[i] = left[i-1]
		} else {
			left[i] = input[i]
		}
	}
	for i := len(input) - 2; i >= 0; i-- {
		if right[i+1] > input[i] {
			right[i] = right[i+1]
		} else {
			right[i] = input[i]
		}
	}
	fmt.Println(input)
	fmt.Println(left)
	fmt.Println(right)
	sum := 0
	for i := 0; i < len(input); i++ {
		minLR := 0
		if left[i] <= right[i] {
			minLR = left[i]
		} else {
			minLR = right[i]
		}
		sum += minLR - input[i]
	}
	fmt.Println(sum)
}
