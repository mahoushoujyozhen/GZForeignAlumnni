package util

import "math"

func Float64ToInt64(num float64) int64 {
	return int64(num * math.Pow10(0))
}

func Float64ToInt64Array(nums []float64) []int64 {
	var transNum []int64
	for i := 0; i < len(nums); i++ {
		transNum = append(transNum, Float64ToInt64(nums[i]))
	}
	return transNum
}
