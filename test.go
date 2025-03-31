package main

import (
	"fmt"
	"math"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	s, e := 1, max(bloomDay)
	for s < e {
		mid := s + (e-s)/2
		if predicate(bloomDay, m, k, mid) {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return s
}

func predicate(bloomDay []int, m, k, mid int) bool {
	bouqe := 0
	for i := 0; i < len(bloomDay); {
		if bloomDay[i] >= mid {
			cnt := 1
			j := 1
			for j <= k {
				if bloomDay[i+j] >= mid {
					cnt++
					j++
				} else {
					break
				}
			}
			if cnt == k {
				bouqe++
			}
			i += j
		}
	}

	return bouqe >= m
}

func max(arr []int) int {
	large := math.MinInt32
	for _, a := range arr {
		if a > large {
			large = a
		}
	}
	return large
}

func main() {
	s := "aaabc"
	freq := make([]int, 256)
	for _, v := range s {
		freq[v]++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})
	fmt.Println(freq)
	// bloomDay := []int{1, 10, 3, 10, 2}
	// m := 3
	// k := 1
	// fmt.Println(minDays(bloomDay, m, k))
}
