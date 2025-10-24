package main

// Runs a closure function with the map's values and returns true if
// any of the value results true with that func.
func mapAnyFunction[K comparable, V any](m map[K]V, f func(V) bool) bool {
	for _, v := range m {
		if f(v) {
			return true
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[byte]int, 9)
	collumnMap := make(map[byte]int, 9)
	var curDigit byte
	for y := range 9 {
		for x := range 9 {
			curDigit = board[y][x]
			if curDigit != '.' {
				rowMap[curDigit]++
			}
			curDigit = board[x][y]
			if curDigit != '.' {
				collumnMap[curDigit]++
			}
		}
		if mapAnyFunction(rowMap, func(val int) bool {
			return val > 1
		}) || mapAnyFunction(collumnMap, func(val int) bool {
			return val > 1
		}) {
			return false
		}
		clear(rowMap)
		clear(collumnMap)
	}
	gridMap := make(map[byte]int, 9)
	for y := 1; y < 9; y += 3 {
		for x := 1; x < 9; x += 3 {
			for a := y - 1; a < y+2; a++ {
				for b := x - 1; b < x+2; b++ {
					curDigit = board[a][b]
					if curDigit == '.' {
						continue
					}
					gridMap[curDigit]++
				}
			}
			if mapAnyFunction(gridMap, func(val int) bool {
				return val > 1
			}) {
				return false
			}
			clear(gridMap)
		}
	}
	return true
}
