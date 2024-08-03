package main

import (
	"fmt"
)

func main() {
	min := 5
	max := 10
	result := rangeSlice(min, max)
	fmt.Println(result)
}

func rangeSlice(min int, max int) []int {
	if min >= max {
		return nil
	}

	length := max - min
	result := []int{}

	for i := 0; i < length; i++ {
		result = append(result, 0)
	}

	for i := range result {
		result[i] = min + i
	}

	return result
}
