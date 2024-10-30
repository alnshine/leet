package maximum_number_of_moves_in_a_grid

var (
	m, n int
	memo [][]int
)

func maxMoves(grid [][]int) int {
	m = len(grid)
	n = len(grid[0])
	memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var res int
	for i := 0; i < m; i++ {
		cur := rec(i, 0, grid)
		if cur > res {
			res = cur
		}
	}
	return res
}

func rec(row, col int, grid [][]int) int {
	if memo[row][col] != -1 {
		return memo[row][col]
	}

	var newRes int
	if row-1 >= 0 && col+1 < n && grid[row][col] < grid[row-1][col+1] {
		newRes = max(newRes, rec(row-1, col+1, grid)+1)
	}
	if col+1 < n && grid[row][col] < grid[row][col+1] {
		newRes = max(newRes, rec(row, col+1, grid)+1)
	}
	if row+1 < m && col+1 < n && grid[row][col] < grid[row+1][col+1] {
		newRes = max(newRes, rec(row+1, col+1, grid)+1)
	}

	memo[row][col] = newRes
	return newRes
}

func Solve(grid [][]int) int {
	return maxMoves(grid)
}
