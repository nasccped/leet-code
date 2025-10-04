package main

import (
	"bytes"
	"strconv"
)

const (
	MIN = -2147483648
	MAX = 2147483647
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var buf bytes.Buffer
	var b byte
	for i := 0; i < len(s); i++ {
		b = s[i]
		if b == ' ' && buf.Len() == 0 {
			continue
		} else if (b == '-' || b == '+') && buf.Len() == 0 {
			buf.WriteByte(b)
			continue
		}
		if b < '0' || b > '9' {
			break
		}
		buf.WriteByte(b)
	}
	asStr := buf.String()
	if len(asStr) == 0 {
		return 0
	} else if len(asStr) == 1 && (asStr[0] == '-' || asStr[0] == '+') {
		return 0
	}
	asInt, err := strconv.Atoi(asStr)
	if err != nil {
		if asStr[0] == '-' {
			return MIN
		} else {
			return MAX
		}
	}
	if asInt < MIN {
		return MIN
	} else if asInt > MAX {
		return MAX
	} else {
		return asInt
	}
}
