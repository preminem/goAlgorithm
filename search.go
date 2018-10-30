package main

import (
	"fmt"
)

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(sequenceSearch(values, 80))
}

func sequenceSearch(values []int, value int) int {
	for i := 0; i < len(values); i++ {
		if values[i] == value {
			return i
		}
	}
	return -1
}
