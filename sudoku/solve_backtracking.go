package sudoku

func solveBacktracking(maze [][]int) bool {
	for i := range 9 {
		for j := range 9 {
			if maze[i][j] != 0 {
				continue
			}

			for k := 1; k <= 9; k++ {
				if validNum(maze, i, j, k) {
					maze[i][j] = k
					if solveBacktracking(maze) {
						return true
					}
					maze[i][j] = 0
				}
			}

			return false
		}
	}

	return true
}

func validNum(maze [][]int, row, cell, num int) bool {
	for i := range 9 {
		if maze[row][i] == num || maze[i][cell] == num {
			return false
		}
	}

	startRow, startCell := (row/3)*3, (cell/3)*3
	for i := range 3 {
		for j := range 3 {
			if maze[startRow+i][startCell+j] == num {
				return false
			}
		}
	}

	return true
}
