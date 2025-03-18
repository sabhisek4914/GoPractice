package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
concat []string in such manner that they are the smallest lexographically.
[a,ab,aba] -> aabaab
*/

func main() {
	str := []string{"a", "ab", "aba"}
	sort.Slice(str, func(i, j int) bool {
		return str[i]+str[j] < str[j]+str[i]
	})
	fmt.Println(strings.Join(str, ""))
}
