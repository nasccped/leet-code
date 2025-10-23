package main

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if nums[i] == target {
			result[0] = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0 && nums[i] >= target; i-- {
		if nums[i] == target {
			result[1] = i
			break
		}
	}
	return result
}
