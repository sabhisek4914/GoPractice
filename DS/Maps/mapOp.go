// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// [['A', 'Laptop'], ['A', 'Laptop'], ['A', 'Mouse'], ['B', 'Laptop'], ['A','Headset'], ['B','Headset'] ]
// Output #1:
// [['A',4], ['B',2]]
// Output #2:
// [['A','Laptop',2], ['A', 'Mouse',1], ['A','Headset',1], ['B','Laptop',1], ['B','Headset',1]]
// Output #3:
// [['A','B',['Laptop','Headset']]]

func main() {
	input := [][]string{{"A", "Laptop"}, {"A", "Laptop"}, {"A", "Mouse"}, {"B", "Laptop"}, {"A", "Headset"}, {"B", "Headset"}}
	fmt.Println(input)
	mp := make(map[string]int)
	for i := 0; i < len(input); i++ {
		x := mp[input[i][0]]
		x++
		mp[input[i][0]] = x
	}
	fmt.Println(mp)
	op1 := make([][]string, 0)
	for key, v := range mp {
		arr := make([]string, 0)
		arr = append(arr, key)
		arr = append(arr, strconv.Itoa(v))
		op1 = append(op1, arr)
	}
	fmt.Println(op1)

	//   mp2:=make(map[string]interface{})
	//   for i:=0;i<len(input);i++{
	//       key:=input[i][0]
	//       label:=input[i][1]
	//       if v,ok:=mp[key];ok{
	//           if v1,ok2:=mp[key].(map[string]interface{})[label].(int);ok{
	//               v1++
	//               mp[key].(map[string]interface{})[label]=v1
	//           }
	//       }
	//   }
	//     fmt.Println(mp2)
	mp3 := make(map[string][]string)
	for i := 0; i < len(input); i++ {
		key := input[i][1]
		x := input[i][0]
		arr := make([]string, 0)
		arr = mp3[key]
		arr = append(arr, x)
		mp3[key] = arr
	}
	fmt.Println(mp3)

	op3 := make([][]interface{}, 0)
	row := make([]interface{}, 0)
	row = append(row, "A")
	row = append(row, "B")
	op3 = append(op3, row)
	// subSlice := []interface{}{"Laptop", "Headset"}
	// x := op3[0]
	// x = append(x, subSlice)
	// op3[0] = x
	// fmt.Println(op3[0][2])

	for k, v := range mp3 {
		a, b := false, false
		for _, arrV := range v {
			if arrV == "A" {
				a = true
			}
			if arrV == "B" {
				b = true
			}
		}
		if a && b {
			fmt.Println(reflect.TypeOf(op3[0][2]))
			if op3[0][2] != nil {
				y := op3[0][2].([]interface{})
				y = append(y, k)
				op3[0][2] = y
			} else {
				y := make([]interface{}, 0)
				y = append(y, k)
				op3[0][2] = y
			}
		}
	}
	fmt.Println(op3)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	input := map[string][]string{
// 		"Headset": {"A", "B"},
// 		"Laptop":  {"A", "A", "B"},
// 		"Mouse":   {"A"},
// 	}

// 	// Create a map to store categories
// 	categoryMap := make(map[string]bool)

// 	// Iterate through the input map
// 	for category := range input {
// 		categoryMap[category] = true
// 	}

// 	// Create an array to store unique categories
// 	var uniqueCategories []string
// 	for category := range categoryMap {
// 		uniqueCategories = append(uniqueCategories, category)
// 	}

// 	// Create the output array
// 	var output [][]interface{}
// 	var categoryNames []string

// 	for _, category := range uniqueCategories {
// 		categoryNames = append(categoryNames, category)

// 		for name, categories := range input {
// 			if contains(categories, category) {
// 				categoryNames = append(categoryNames, name)
// 			}
// 		}

// 		output = append(output, []interface{}{categoryNames[1], categoryNames[2], categoryNames[0]})
// 		categoryNames = []string{}
// 	}

// 	// Print the output
// 	fmt.Println(output)
// }

// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	input := [][]string{
// 		{"A", "Laptop"},
// 		{"A", "Laptop"},
// 		{"A", "Mouse"},
// 		{"B", "Laptop"},
// 		{"A", "Headset"},
// 		{"B", "Headset"},
// 	}

// 	// Create a map to store counts
// 	counts := make(map[string]int)

// 	// Iterate over the input array and count the combinations
// 	for _, item := range input {
// 		key := item[0] + "," + item[1]
// 		counts[key]++
// 	}

// 	// Convert the map into the desired output format
// 	var output [][]interface{}
// 	for key, count := range counts {
// 		parts := splitKey(key)
// 		output = append(output, append(parts, count))
// 	}

// 	// Print the output
// 	fmt.Println(output)
// }

// // Helper function to split the key into parts
// func splitKey(key string) []interface{} {
// 	parts := make([]interface{}, 0)
// 	for _, part := range split(key, ',') {
// 		parts = append(parts, part)
// 	}
// 	return parts
// }

// // Helper function to split a string into parts
// func split(s, sep string) []string {
// 	parts := strings.Split(s, sep)
// 	return parts
// }
