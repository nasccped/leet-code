package main

func numIslands(grid [][]byte) int {
	var (
		count, m, n int = 0, 0, 0
		isIsland bool = false
		explore func(y, x int)
	)
	explore = func(y, x int) {
		if (y < 0 || y >= m) || (x < 0 || x >= n) || grid[y][x] == '0' {
			return
		}
		isIsland = true
		grid[y][x] = '0'
		explore(y-1, x)
		explore(y+1, x)
		explore(y, x-1)
		explore(y, x+1)
	}
	m = len(grid)
	n = len(grid[0])
	for y := range m {
		for x := range n {
			explore(y, x)
			if isIsland {
				count++
				isIsland = false
			}
		}
	}
	return count
}
