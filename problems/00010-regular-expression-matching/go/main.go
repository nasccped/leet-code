package main

import (
	"bytes"
	"regexp"
)

// Builds the suitable regex pattern
func buildRegex(pattern string) string {
	var buf bytes.Buffer
	buf.WriteString("^(")
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '.':
			buf.WriteString("[a-z]")
		case '*':
			latest := buf.Bytes()[buf.Len()-1]
			if latest != ']' {
				buf.Truncate(buf.Len() - 1)
				buf.Write([]byte{'(', latest, ')'})
			}
			buf.WriteByte('*')
		default:
			buf.WriteByte(pattern[i])
		}
	}
	buf.WriteString(")$")
	return buf.String()
}

func isMatch(s string, p string) bool {
	if success, err := regexp.MatchString(buildRegex(p), s); err != nil {
		return false
	} else {
		return success
	}
}
