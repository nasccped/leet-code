package main

// Coordinate type abstraction
type Coord struct {
	y, x int
}

func numIslands(grid [][]byte) int {
	var (
		count   int                = 0
		set     map[Coord]struct{} = map[Coord]struct{}{}
		explore func(y, x, counter int) int
	)
	explore = func(y, x, counter int) int {
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
			return counter
		}
		c := Coord{y, x}
		if _, exists := set[c]; exists || grid[y][x] == '0' {
			return counter
		}
		set[c] = struct{}{}
		counter++
		return explore(y-1, x, counter) + explore(y+1, x, counter) + explore(y, x-1, counter) + explore(y, x+1, counter)
	}
	for y, row := range grid {
		for x, slot := range row {
			if slot == '0' {
				continue
			}
			if explore(y, x, 0) > 0 {
				count++
			}
		}
	}
	return count
}
