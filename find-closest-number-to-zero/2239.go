package find_closest_number_to_zero

import (
	"math"
)

func findClosestNumber(nums []int) int {
	var res int
	minN := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if val < 0 {
			val *= -1
		}
		if val < minN {
			minN = val
			res = nums[i]
		} else if val == minN && res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

func Solve(arr []int) int {
	return findClosestNumber(arr)
}
