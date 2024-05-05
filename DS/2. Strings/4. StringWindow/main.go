package main

import (
	"fmt"
	"math"
)

//s1:="hello_world"
//s2:=lol
//op:="llo"  -->> i.e. the smallest

//Brute force find all substrings -> filter the string that contains "lol" in it -> get the smallest out of it.

func main() {
	s1 := "helloworld"
	s2 := "lol"
	small := math.MaxInt32
	f_s1 := make([]int, 256)
	f_s2 := make([]int, 256)
	for i := 0; i < len(s2); i++ {
		f_s2[s2[i]]++
		// fmt.Println(int(s2[i]))
	}
	// fmt.Println(f_s2)

	start := 0
	op := ""
	cnt := 0
	for i := 0; i < len(s1); i++ {
		c := s1[i]
		// ch := string(s1[i])
		f_s1[c]++
		if f_s2[c] > 0 {
			cnt++
		}
		if cnt >= len(s2) {
			// fmt.Println(f_s1[108], f_s1[111])
			// fmt.Println("OP: ", s1[start:i+1], cnt)
			for f_s2[s1[start]] == 0 || f_s1[s1[start]] > f_s2[s1[start]] {
				f_s1[s1[start]]--
				start++
			}
			window_size := i - start + 1
			// fmt.Println(s1[:i+1])
			// fmt.Println(s1[start:window_size])
			if window_size < small {
				small = window_size
				op = s1[start : start+small]
			}
		}
	}
	fmt.Println(op)
}
