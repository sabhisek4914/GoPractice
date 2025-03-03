package main

import "fmt"

/*
Housing Problem:- Near road side there are some plots of variable areas sequencially
You want ot buy k acers of land to build a house. Find all the plots whose sum=k
Example:
1, 3, 2, 1, 4, 1, 3
0, 1, 2, 3, 4, 5, 6
op: 2,5 and 4,6

2pointer -> O(n)

Prefix sum and search
	1, 3, 2, 1, 4, 1, 3 -> 1, 4, 6, 7, 11, 12, 15 -> O(n)
	for every element i -> arr[i]+k= desired element present in prefixSum[] or not. -> O(n)
		search that desiredElelemnt usingbinary search -> logn
TimeCOmplexity -> O(nlogn )
*/

func main() {
	arr := []int{1, 3, 2, 1, 4, 1, 3}
	BuyLand(arr, 8)
}

func BuyLand(arr []int, k int) {
	cSum := 0
	i, j := 0, 0
	for j < len(arr) {
		cSum += arr[j]
		for cSum > k && i < j {
			cSum -= arr[i]
			i++
		}
		if cSum == k {
			fmt.Println(i, j)
		}
		j++
	}
}
