package main

import "fmt"

/*
54. Spiral Matrix -> https://leetcode.com/problems/spiral-matrix/description/
Given an m x n matrix, return all elements of the matrix in spiral order.
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Time complexity: O(m×n)
We visit each element of the matrix once.

Space complexity: O(m×n)
Storing the output list for the response -> m*n no. of elements we are storing in 1D []
*/

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	op := spiralOrder(matrix)
	fmt.Println(op)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	op := make([]int, 0)

	nextDirection := map[string]string{
		"top":    "right",
		"right":  "bottom",
		"bottom": "left",
		"left":   "top",
	}
	direction := "top"
	rTop, rBot := 0, m-1
	cLeft, cRight := 0, n-1
	for len(op) < m*n {
		if direction == "top" {
			for i := cLeft; i <= cRight; i++ {
				op = append(op, matrix[rTop][i])
			}
			rTop++
		} else if direction == "bottom" {
			for i := cRight; i >= cLeft; i-- {
				op = append(op, matrix[rBot][i])
			}
			rBot--
		} else if direction == "right" {
			for i := rTop; i <= rBot; i++ {
				op = append(op, matrix[i][cRight])
			}
			cRight--
		} else if direction == "left" {
			for i := rBot; i >= rTop; i-- {
				op = append(op, matrix[i][cLeft])
			}
			cLeft++
		}
		direction = nextDirection[direction]
	}
	return op
}
