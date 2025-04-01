package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func clampStringToInt32(s string) int {
	// Trim whitespace
	s = strings.TrimSpace(s)

	// Convert to integer
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// Clamp the number within 32-bit signed integer range
	if num < math.MinInt32 {
		return math.MinInt32
	} else if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return num
}

func main() {
	fmt.Println(clampStringToInt32("2147483648"))  // Output: 2147483647 (Clamped)
	fmt.Println(clampStringToInt32("-2147483649")) // Output: -2147483648 (Clamped)
	fmt.Println(clampStringToInt32("12345"))       // Output: 12345 (Within range)
	fmt.Println(clampStringToInt32("1337.456"))    // Output: 0 (Invalid case)
}
