package main

import (
	"fmt"
	"math"
)

/*
1901. Find a Peak Element II -> https://leetcode.com/problems/find-a-peak-element-ii/description/
A peak element in a 2D grid is an element that is strictly greater than all of its adjacent neighbors to the left, right, top, and bottom.
Given a 0-indexed m x n matrix mat where no two adjacent cells are equal, find any peak element mat[i][j] and return the length 2 array [i,j].
You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.
You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.

Input: mat = [[10,20,15],[21,30,14],[7,16,32]]
Output: [1,1]
Explanation: Both 30 and 32 are peak elements so [1,1] and [2,2] are both acceptable answers.

Approach:-> FindPeak1 logic u use for cloumns, to find max row use O(n) -> large ele
then mat[maxRow][m] >= mat[maxRow][m+1]
*/

func main() {
	mat := [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}
	fmt.Println(findPeakGrid(mat))
}

func findPeakGrid(mat [][]int) []int {
	s, e := 0, len(mat[0])-1
	for s < e {
		m := s + (e-s)/2
		if predicate(mat, m) {
			e = m
		} else {
			s = m + 1
		}
	}
	maxRow := maxRow(mat, s)
	return []int{maxRow, s}
}

func predicate(mat [][]int, m int) bool {
	maxRow := maxRow(mat, m)
	if m+1 <= len(mat[0])-1 {
		fmt.Println(maxRow, m)
		return mat[maxRow][m] >= mat[maxRow][m+1]
	}
	return false
}
func maxRow(mat [][]int, m int) int {
	large := math.MinInt32
	ele := 0
	for i := 0; i < len(mat); i++ {
		if mat[i][m] > large {
			large = mat[i][m]
			ele = i
		}
	}
	return ele
}
