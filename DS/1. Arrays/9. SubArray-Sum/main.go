package main

import (
	"fmt"
	"math"
)

// input:={-1,2,3,4,-2,6,-8,3}
// op:= 13 => {2,3,4,-2,6}

func main() {
	input := []int{-1, 2, 3, 4, -2, 6, -8, 3}
	large := math.MinInt64
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
		if sum < 0 {
			sum = 0
		}
		if sum > large {
			large = sum
		}
	}
	fmt.Println(large)
}
