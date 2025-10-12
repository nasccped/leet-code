package main

import "slices"

func removeDuplicates(nums []int) int {
	newSlice := []int{}
	for _, n := range nums {
		if !slices.Contains(newSlice, n) {
			newSlice = append(newSlice, n)
		}
	}
	for i, n := range newSlice {
		nums[i] = int(n)
	}
	return len(newSlice)
}
