package main

import "slices"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	indexes := []int{}
	for i := 0; i < (len(s) + numRows); i += numRows*2 - 2 {
		indexes = append(indexes, i)
	}
	for i := 0; i < len(indexes); i++ {
		otherInd := indexes[i]
		otherInd -= 1
		if !(otherInd < 0 || slices.Contains(indexes, otherInd)) {
			indexes = append(indexes, otherInd)
		}
		otherInd += 2
		if !(otherInd > len(s) || slices.Contains(indexes, otherInd)) {
			indexes = append(indexes, otherInd)
		}
	}
	result := []byte{}
	for _, ind := range indexes {
		if ind >= len(s) {
			continue
		}
		result = append(result, s[ind])
	}
	return string(result)
}
