package main

import "strconv"

// Solving with string convertion (better)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	runes := []rune(strconv.Itoa(x))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

// Solving without string convertion
func isPalindromeNoString(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	accum := 0
	for i := 1; x/i > 0; i *= 10 {
		accum *= 10
		accum += (x / i) % 10
	}
	return accum == x
}
