package main

import "fmt"

/*
73. Set Matrix Zeroes -> https://leetcode.com/problems/set-matrix-zeroes/description/
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place. (Don't create a ny copy of original matrix)

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?

For O(1) solution:
1. m*n logic iterate to find a[i][j]==0
	then only set it's i,0 and 0,j value as 0
		a[i][0]=0
		a[0][j]=0
2. m*n logic start from 1(bcz if 0,0 == 0 then it will set all values to 0) -> if a[i][0] == 0 || a[0][j] == 0 -> a[i][j]=0
3. see the logic

1,	2,	0,	4
5,	0,	6,	7
1,	8,	9,	2

1,	0,	0,	4
0,	0,	6,	7
1,	8,	9,	2

1,	0,	0,	4
0,	0,	0,	0
1,	0,	0,	2

0,	0,	0,	0
0,	0,	0,	0
1,	0,	0,	2
*/

func setZeroesSpaceOptimal(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	first_column := false
	first_row := false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			first_row = true
			break
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			first_column = true
			break
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	//** imp i=1,j=1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if first_row {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if first_column {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	row := make([]bool, n)
	col := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	arr := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(arr)
	fmt.Println(arr)
	arr2 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroesSpaceOptimal(arr2)
	fmt.Println(arr2)
}
