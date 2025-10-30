package main

import "slices"

// Operates the sum oper. with the given slice.
func sumSlice(slc []int) int {
	result := 0
	for _, n := range slc {
		result += n
	}
	return result
}

// Checks if two slices are equals (no sort required).
func sliceUnsortedEquals[T comparable](s1, s2 []T) bool {
	m := map[T]int{}
	for _, item := range s1 {
		m[item]++
	}
	for _, item := range s2 {
		m[item]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var backtrack func(slc []int)
	backtrack = func(slc []int) {
		sum := sumSlice(slc)
		if sum == target && !slices.ContainsFunc(result, func(s []int) bool {
			return sliceUnsortedEquals(s, slc)
		}) {
			result = append(result, slices.Clone(slc))
		}
		for _, c := range candidates {
			if sum+c <= target {
				slc = append(slc, c)
				backtrack(slc)
				slc = slc[:len(slc)-1]
			}
		}
	}
	for _, c := range candidates {
		backtrack([]int{c})
	}
	return result
}
