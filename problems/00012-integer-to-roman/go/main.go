package main

import "bytes"

// Returns a roman repr. pattern since they're all the same.
//
// It uses 'a', 'b' and 'c' bytes to represent the pattern:
//
// 'a' means I, X, C or M
//
// 'b' means V, L or D
//
// 'c' means the same of 'a' but with 1 offset (i.e. when 'a' is I, 'c' is X...)
//
// This function will return the suitable pattern only if the num is within 1..9
// inclusive range, otherwise, it'll return an empty slice. The returned value
// should be handle and the 'abc' symbols replaced by the suitable ones.
func romanPattern(num int) []byte {
	switch num {
	case 1, 2, 3:
		return bytes.Repeat([]byte{'a'}, num)
	case 4:
		return []byte{'a', 'b'}
	case 5:
		return []byte{'b'}
	case 6, 7, 8:
		ret := []byte{'b'}
		ret = append(ret, bytes.Repeat([]byte{'a'}, num-5)...)
		return ret
	case 9:
		return []byte{'a', 'c'}
	default:
		return []byte{}
	}

}

func intToRoman(num int) string {
	romans := []byte("IVXLCDM")
	symbolInd := 1
	var a, b, c byte
	romansString := []string{}
	var n int
	// Get the roman symbol by the `ind`. returns '0' when out of range.
	byteOrNil := func(ind int) byte {
		if ind >= len(romans) || ind < 0 {
			return '0'
		}
		return romans[ind]
	}
	for ; num > 0; num /= 10 {
		a, b, c = byteOrNil(symbolInd-1), byteOrNil(symbolInd), byteOrNil(symbolInd+1)
		n = num % 10
		if n == 0 {
			symbolInd += 2
			continue
		}
		bt := romanPattern(n)
		bt = bytes.ReplaceAll(bt, []byte{'a'}, []byte{a})
		bt = bytes.ReplaceAll(bt, []byte{'b'}, []byte{b})
		bt = bytes.ReplaceAll(bt, []byte{'c'}, []byte{c})
		romansString = append(romansString, string(bt))
		symbolInd += 2
	}
	var buf bytes.Buffer
	for i := len(romansString) - 1; i >= 0; i-- {
		buf.WriteString(romansString[i])
	}
	return buf.String()
}
