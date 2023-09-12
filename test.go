package main

import (
	"fmt"
	"strings"
	// "strings"
)

//input = "Subham Abhisek"

func main() {
	input := "Subham Abhisek"
	fmt.Println(strings.ToLower(input))
	lower := strings.ToLower(input)
	fmt.Println("Lower", lower)
	mp := make(map[string]int)
	for i := 0; i < len(lower); i++ {
		if string(lower[i]) != " " && string(lower[i]) != "a" && string(lower[i]) != "e" && string(lower[i]) != "i" && string(lower[i]) != "o" && string(lower[i]) != "u" {
			mp[string(lower[i])]++
		}
	}
	fmt.Println(mp)
	mpCheck := make(map[string]bool)
	for i := 0; i < len(lower); i++ {
		if mp[string(lower[i])] >= 1 && !mpCheck[string(lower[i])] {
			fmt.Printf("%s : %d ", string(lower[i]), mp[string(lower[i])])
			mpCheck[string(lower[i])] = true
		}
	}
}
