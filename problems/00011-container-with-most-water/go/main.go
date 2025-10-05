package main

// Turns an integer into its absolute value (math lib provides only float64 abs)
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxArea(height []int) int {
	leftInd, rightInd := 0, len(height)-1
	maxi := 0
	for leftInd < rightInd {
		y := min(height[leftInd], height[rightInd])
		x := absInt(leftInd - rightInd)
		if temp := x * y; temp > maxi {
			maxi = temp
		}
		if height[leftInd] < height[rightInd] {
			leftInd++
		} else {
			rightInd--
		}
	}
	return maxi
}
