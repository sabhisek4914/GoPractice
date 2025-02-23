package main

import (
	"fmt"
	"sort"
)

// https://www.geeksforgeeks.org/minimum-number-swaps-required-sort-array/
//2366 https://leetcode.com/problems/minimum-replacements-to-sort-the-array/description/ need to try
// [10,11,5,4,3,2,1]
// sorted [1,2,3,4,5,10,11]

type Pair struct {
	First  int
	Second int
}

func main() {
	arr := []int{10, 11, 5, 4, 3, 2, 1, 9}
	mp := make([]map[int]int, 0)
	var pr []Pair
	for i := 0; i < len(arr); i++ {
		m := make(map[int]int)
		m[arr[i]] = i
		mp = append(mp, m)

		p := Pair{arr[i], i}
		pr = append(pr, p)
	}
	fmt.Println(mp)
	fmt.Println(pr)
	sort.Slice(mp, func(a, b int) bool {
		key1 := getKeys(mp[a])
		key2 := getKeys(mp[b])
		return key1[0] < key2[0]
	})
	sort.Slice(pr, func(a, b int) bool {
		return pr[a].First < pr[b].First
	})
	fmt.Println(mp)
	fmt.Println(pr)
	visited := make([]bool, 0)
	for i := 0; i < len(pr); i++ {
		x := false
		visited = append(visited, x)
	}

	cnt := 0
	for i := 0; i < len(pr); i++ {
		old := pr[i].Second
		if visited[i] == true || old == i {
			continue
		}
		cycle := 0
		node := i
		for visited[node] == false {
			visited[node] = true
			cycle++
			node = pr[node].Second
		}
		if cycle > 0 {
			cnt += cycle - 1
		}
	}
	fmt.Println(cnt)
}
func getKeys(m map[int]int) []int {
	op := make([]int, 0)
	for k := range m {
		op = append(op, k)
	}
	return op
}
