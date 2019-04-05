package pony_cross_river

import "math"

func PonyCrossRiver(matrix [][]int) int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return 0
	}
	res := math.MaxInt32
	for i, _ := range matrix[0] {
		sum := ponyCrossRiverCore(matrix, 0, i)
		if sum < res {
			res = sum
		}
	}
	return res
}

func ponyCrossRiverCore(matrix [][]int, i int, j int) int {
	current := matrix[i][j]
	var stepA, stepB, stepC, stepD int
	if j-2 >= 0 && i+1 < len(matrix) {
		stepA = ponyCrossRiverCore(matrix, i+1, j-2)
	}
	if j+2 < len(matrix[0]) && i+1 < len(matrix) {
		stepB = ponyCrossRiverCore(matrix, i+1, j+2)
	}
	if j-1 >= 0 && i+2 < len(matrix) {
		stepC = ponyCrossRiverCore(matrix, i+2, j-1)
	}
	if j+1 < len(matrix[0]) && i+2 < len(matrix) {
		stepD = ponyCrossRiverCore(matrix, i+2, j+1)
	}
	return current + min([]int{stepA, stepB, stepC, stepD})
}

func min(arr []int) int {
	minVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if minVal > arr[i] {
			minVal = arr[i]
		}
	}
	return minVal
}
