package main

import "math"

func increasingTriplet(nums []int) bool {
	var first_num = math.MaxInt64
	var second_num = math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if nums[i] <= first_num {
			first_num = nums[i]
		} else if nums[i] <= second_num {
			second_num = nums[i]
		} else {
			return true
		}
	}
	return false
}
