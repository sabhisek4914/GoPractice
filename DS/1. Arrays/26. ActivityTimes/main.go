package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Find max no. of activities can be done, given pair of start and end time of each activity
(7,9),(0,10),(4,5),(8,9),(4,10),(5,7)
OP: 3
(4,5),(5,7),(7,9)

approach, sort by 2nd vals
*/

type Pair struct {
	First  int
	Second int
}

func main() {
	inp := []Pair{
		{7, 9}, {0, 10}, {4, 5}, {8, 9}, {4, 10}, {5, 7},
	}
	sort.Slice(inp, func(i, j int) bool {
		return inp[i].Second < inp[j].Second
	})
	fmt.Println(inp)
	previous := inp[0].Second
	large := math.MinInt32
	cnt := 1
	for i := 0; i < len(inp); i++ {
		if inp[i].First >= previous {
			cnt++
			previous = inp[i].Second
		} else {
			// fmt.Println(cnt)
			if cnt > large {
				large = cnt
			}
			cnt = 1
		}
	}
	fmt.Println(large)
}
