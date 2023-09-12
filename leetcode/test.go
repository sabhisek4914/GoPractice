package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Hello %T %v", 10, 10)
	var x = fmt.Sprintf("Hello %T %v", 10, 10)
	fmt.Println(x)
	y := []int{19, 10, 11}
	sort.Ints(y)
	fmt.Println(y)
}
