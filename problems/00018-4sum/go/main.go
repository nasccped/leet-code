package main

import (
	"sort"
)

// Returns the sum of an integer slice.
func sliceSum(slice []int) int {
	accum := 0
	for _, item := range slice {
		accum += item
	}
	return accum
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for d := len(nums) - 1; d >= a+3; d-- {
			if d < len(nums)-1 && nums[d] == nums[d+1] {
				continue
			}
			b, c := a+1, d-1
			for b < c {
				curSlice := []int{nums[a], nums[b], nums[c], nums[d]}
				if temp := sliceSum(curSlice); temp == target {
					result = append(result, curSlice)
					for b < c && nums[b] == nums[b+1] {
						b++
					}
					for b < c && nums[c] == nums[c-1] {
						c--
					}
					b++
					c--
				} else if temp < target {
					b++
				} else {
					c--
				}
			}
		}
	}
	return result
}
