package main

const (
	MIN = -2147483648
	MAX = 2147483647
)

func reverse(x int) int {
	value := 0
	is_neg := false
	if x < 0 {
		x *= -1
		is_neg = true
	}
	for x > 0 {
		value *= 10
		value += x % 10
		x /= 10
	}
	if is_neg {
		value *= -1
	}
	if value < MIN || value > MAX {
		return 0
	}
	return value
}
