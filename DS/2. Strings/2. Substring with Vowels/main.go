package main

import (
	"fmt"
	"strings"
)

// Brute Force
func main() {
	str := "mnarenbohidumnarhidgem"
	// str := "abc"
	op := make([]string, 0)
	for i := 0; i < len(str); i++ {
		for j := i + i; j <= len(str); j++ {
			a := str[i:j]
			if strings.Contains(str[i:j], "a") && strings.Contains(str[i:j], "e") && strings.Contains(str[i:j], "i") && strings.Contains(str[i:j], "o") && strings.Contains(str[i:j], "u") {
				op = append(op, str[i:j])
			}
			fmt.Println(a)
		}
	}
	fmt.Println(len(op))
}

// func main() {
// 	str := "mnarenbohidumnarhidgem"
// 	i := 0
// 	j := 0
// 	op := make([]string, 0)
// 	for j < len(str) {
// 		//window:=str[i:j]
// 		for strings.Contains(str[i:j], "a") && strings.Contains(str[i:j], "e") && strings.Contains(str[i:j], "i") && strings.Contains(str[i:j], "o") && strings.Contains(str[i:j], "u") {
// 			op = append(op, str[i:j])
// 			i++
// 		}
// 		i = 0
// 		j++
// 	}
// 	fmt.Println(len(op))
// }
