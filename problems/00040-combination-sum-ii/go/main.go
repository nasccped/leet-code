package main

import "slices"

var globalMap map[int]int = map[int]int{}

// Check if two slices contains the same elements without ordering
// constraint.
func unorderedEquals(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range globalMap {
		delete(globalMap, k)
	}
	for _, k := range s1 {
		globalMap[k]++
	}
	for _, k := range s2 {
		globalMap[k]--
	}
	for _, v := range globalMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	countMap := [][]int{}
	for _, c := range candidates {
		if temp := slices.IndexFunc(countMap, func(item []int) bool {
			return item[0] == c
		}); temp != -1 {
			countMap[temp][1]++
		} else {
			countMap = append(countMap, []int{c, 1})
		}
	}
	var backtrack func(s []int, curVal int)
	backtrack = func(s []int, curVal int) {
		if curVal == target && !slices.ContainsFunc(result, func(items []int) bool {
			return unorderedEquals(items, s)
		}) {
			result = append(result, slices.Clone(s))
			return
		} else if curVal > target {
			return
		}
		for _, item := range countMap {
			if item[1] == 0 {
				continue
			}
			item[1]--
			s = append(s, item[0])
			backtrack(s, curVal + item[0])
			item[1]++
			s = s[:len(s)-1]
		}
	}
	backtrack([]int{}, 0)
	return result
}
