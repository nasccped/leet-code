package main

// Checks if a byte arr (string) is a palindrome
func isPalindrome(arr []byte) bool {
	if len(arr) <= 1 {
		return true
	}
	left, right := 0, len(arr)-1
	for ; left < right; left, right = left+1, right-1 {
		if arr[left] != arr[right] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	bytes := []byte(s)
	var cur_slice []byte
	var result = ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j >= i; j-- {
			cur_slice = bytes[i:j]
			if isPalindrome(cur_slice) && len(cur_slice) > len(result) {
				result = string(cur_slice)
			}
		}
	}
	return result
}
