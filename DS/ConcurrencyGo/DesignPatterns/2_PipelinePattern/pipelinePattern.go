package main

import "fmt"

/*
The pipeline pattern structures a series of processing stages, with each stage executed concurrently.
Data flows through these stages sequentially, allowing efficient data transformation and processing.
Application:
1. Data processing in ETL (Extract, Transform, Load) pipelines.
2. Image processing pipelines in multimedia applications.
*/

func main() {
	//create initail channelw ith some data
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))
	for _, d := range data {
		input <- d
	}
	close(input)

	//First Stage of pipeline: Doules the input value
	doubleOutput := make(chan int)
	go func() {
		defer close(doubleOutput)
		for num := range input {
			doubleOutput <- num * 2
		}
	}()

	//2nd Stage of pipeline: Squares the input value
	squaredOutput := make(chan int)
	go func() {
		defer close(squaredOutput)
		for num := range doubleOutput {
			fmt.Println("DoubledOutput: ", num)
			squaredOutput <- num * num
		}
	}()

	//3rd stage of pipeline: Print squared values
	for result := range squaredOutput {
		fmt.Println("SquaredOutput: ", result)
	}
}
