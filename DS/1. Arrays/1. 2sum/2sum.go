package main

import "fmt"

// {-1,0,1,2,3,5,6,8} k=8
func printPairs(arr []int, sum int) {
	mp := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		req := sum - arr[i]
		if mp[req] {
			fmt.Println(arr[i], " : ", req)
		}
		mp[arr[i]] = true
	}
	fmt.Println(mp)
}
func main() {
	arr := []int{-1, 0, 1, 2, 3, 5, 6, 8}
	sum := 8
	printPairs(arr, sum)
}
