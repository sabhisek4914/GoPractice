package main

import (
	"fmt"
	"math"
)

//input:= "aabcb"
//op:= "abc"

//Brute Force
// func checkDuplicate(str string) bool {
// 	mp := make(map[string]struct{})
// 	for i := 0; i < len(str); i++ {
// 		mp[string(str[i])] = struct{}{}
// 	}
// 	if len(str) == len(mp) {
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	str := "aabcb"
// 	large := math.MinInt64
// 	op := ""
// 	for i := 0; i < len(str); i++ {
// 		for j := i + 1; j < len(str); j++ {
// 			if checkDuplicate(str[i:j]) {
// 				if len(str[i:j]) > large {
// 					large = len(str[i:j])
// 					op = str[i:j]
// 					fmt.Println(str[i:j])
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(op)
// }

func main() {
	str := "aabcb"
	large := math.MinInt64
	op := ""
	i, j := 0, 0
	window_len := 0
	mp := make(map[string]int)
	mpBool := make(map[string]bool)
	for j < len(str) {
		ch := string(str[j])
		if mpBool[ch] && mp[ch] >= i {
			i = mp[ch] + 1
			window_len = j - i
		}
		mp[ch] = j
		mpBool[ch] = true
		window_len++
		j++
		if window_len > large {
			large = window_len
			op = str[i : i+large]
		}
	}
	fmt.Println(op)
}
