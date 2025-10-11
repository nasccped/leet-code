package main

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(pattern string, opennings, closings int)
	backtrack = func(pattern string, opennings, closings int) {
		if len(pattern) == n*2 {
			result = append(result, pattern)
			return
		}
		if opennings < n {
			backtrack(pattern+"(", opennings+1, closings)
		}
		if closings < opennings {
			backtrack(pattern+")", opennings, closings+1)
		}
	}
	backtrack("", 0, 0)
	return result
}
