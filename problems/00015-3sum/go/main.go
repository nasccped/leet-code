package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	var j, k int
	for i, a := range nums {
		if i >= len(nums)-2 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k = i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] == -a {
				result = append(result, []int{a, nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[j]+nums[k] < -a {
				j++
			} else {
				k--
			}
		}
	}
	return result
}
