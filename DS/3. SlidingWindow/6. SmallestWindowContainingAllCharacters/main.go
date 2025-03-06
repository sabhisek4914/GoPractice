package main

import (
	"fmt"
	"math"
)

/*
SmallesWindow that has all the characters of a substring.
str:="aabcbcdbca"
op:="dbca"
*/

func main() {
	str1 := "aabcbcdbca"
	fmt.Println(smallestWindowSubString(str1))
}

func smallestWindowSubString(s string) string {
	mp := make(map[rune]bool)
	freqS, freqP := [256]int{}, [256]int{}
	// sArr := strings.Split(s, "")
	for _, ch := range s {
		mp[ch] = true
	}
	for k, _ := range mp {
		freqP[int(k)]++
	}
	cnt := 0
	start := 0
	small := math.MaxInt32
	str := ""
	for i := 0; i < len(s); i++ {
		freqS[s[i]]++
		if freqP[s[i]] != 0 && freqS[s[i]] <= freqP[s[i]] {
			cnt++
		}
		// fmt.Println(cnt, start, i)
		if cnt == len(mp) {
			for freqP[s[start]] == 0 || freqS[s[start]] > freqP[s[start]] {
				freqS[s[start]]--
				start++
			}
			fmt.Println(cnt, start, i)
			length := i - start + 1
			if length < small {
				small = length
				str = s[start : i+1]
				fmt.Println(str, cnt, start)
			}
		}
	}
	if small == math.MaxInt32 {
		return ""
	}
	return str
}
