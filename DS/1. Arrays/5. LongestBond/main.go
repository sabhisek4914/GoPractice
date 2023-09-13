package main

import "fmt"

//i/p: [1,9,3,0,18,5,2,4,10,7,12,6]
//op: 8
// 0,1,2,3,4,5,6,7

func main() {
	input := []int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6}
	large := 0
	mp := make(map[int]bool)
	for i := 0; i < len(input); i++ {
		mp[input[i]] = true
	}
	for i := 0; i < len(input); i++ {
		parent := input[i] - 1
		if mp[parent] == false {
			//current input[i] is the parent element
			ctr := 0
			p := input[i]
			for mp[p] {
				ctr++
				p++
			}
			if ctr > large {
				large = ctr
			}
		}
	}
	fmt.Println(large)
}
