package main

import "maps"

func findSubstring(s string, words []string) []int {
	wordCount := map[string]int{}
	syllLen := len(words[0])
	wordLen := len(words) * syllLen
	for _, w := range words {
		wordCount[w]++
	}
	inds := []int{}
	copied := make(map[string]int, len(wordCount))
sIter:
	for i := 0; i <= len(s)-wordLen; i++ {
		maps.Copy(copied, wordCount)
		for j := i; j < i+wordLen; j += syllLen {
			copied[s[j:j+syllLen]]--
			if x := copied[s[j:j+syllLen]]; x < 0 {
				maps.DeleteFunc(copied, func(key string, value int) bool {
					return true
				})
				continue sIter
			}
		}
		inds = append(inds, i)
	}
	return inds
}
