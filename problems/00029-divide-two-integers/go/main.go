package main

const (
	MIN_VALUE = -2147483648
	MAX_VALUE = 2147483647
)

// Returns the absolute value of an integer.
func absInt(val int) int {
	if val < 0 {
		return val * -1
	} else {
		return val
	}
}

func divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}
	isNegative := (dividend < 0) != (divisor < 0)
	dividend, divisor = absInt(dividend), absInt(divisor)
	count := 0
	for dividend >= divisor {
		temp := 0
		for dividend > (divisor << (temp + 1)) {
			temp++
		}
		count += 1 << temp
		dividend -= divisor << temp
	}
	if isNegative {
		if count < MIN_VALUE {
			return MIN_VALUE
		} else {
			return -count
		}
	} else {
		if count > MAX_VALUE {
			return MAX_VALUE
		} else {
			return count
		}
	}
}
