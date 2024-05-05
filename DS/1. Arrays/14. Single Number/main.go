package main

import "fmt"

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Input: nums = [2,2,1]
Output: 1

Input: nums = [4,1,2,1,2]
Output: 4
*/

func main() {

	//Frequency Map
	mp := make(map[int]int)
	arr := []int{4, 1, 2, 1, 2}
	for _, x := range arr {
		mp[x]++
	}
	fmt.Println(mp)
	for k, v := range mp {
		if v == 1 {
			fmt.Println(k)
		}
	}

	//2nd solution
	ans := 0
	for _, x := range arr {
		ans ^= x
	}
	fmt.Println(ans)

}
