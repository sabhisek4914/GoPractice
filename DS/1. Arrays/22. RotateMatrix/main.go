package main

import "fmt"

/*
48. Rotate Image -> https://leetcode.com/problems/rotate-image/description/
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Approach:
To achieve a 90-degree rotation, two key transformations are needed:
1. Vertical Reversal (Flip along the horizontal axis)
2. Transpose (Swap rows and columns)

Combining these steps effectively rotates the matrix. Let’s break it down:
1. Vertical Reversal
In this step, we flip the matrix upside down. For every column, the topmost element swaps with the bottommost element, the second topmost swaps with the second bottommost, and so on.
Visualization: Vertical Reversal
Input Matrix:
1  2  3
4  5  6
7  8  9
After Vertical Reversal:
7  8  9
4  5  6
1  2  3
This aligns the rows correctly for the next transformation.

2. Transpose
Transposing a matrix swaps the rows and columns. In simpler terms, we swap matrix[i][j] with matrix[j][i] for every valid pair of indices above the diagonal.
Visualization: Transpose
Matrix after Vertical Reversal:
7  8  9
4  5  6
1  2  3
After Transposing:
7  4  1
8  5  2
9  6  3
This final transformation achieves the desired clockwise rotation.
*/

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[len(matrix)-1-i][j]
			matrix[len(matrix)-1-i][j] = temp
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
