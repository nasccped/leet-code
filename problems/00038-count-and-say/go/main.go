package main

import "strings"

// Generates a RLE from a given string.
func RLEFromString(s string) string {
	var (
		buf   strings.Builder
		count byte
		b     byte
	)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			b = s[i]
			count = 1
			continue
		}
		if s[i] == s[i-1] {
			count++
			continue
		}
		buf.WriteByte('0' + count)
		buf.WriteByte(b)
		b = s[i]
		count = 1
	}
	buf.WriteByte('0' + count)
	buf.WriteByte(b)
	return buf.String()
}

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		result = RLEFromString(result)
	}
	return result
}
