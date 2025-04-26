package sudoku

import "fmt"

func Solve2(maze [][]int) ([][]int, error) {
	if err := ValidateInput(maze); err != nil {
		return nil, err
	}

	m := makeEmptyMap()
	for i := range 9 {
		for j := range 9 {
			if maze[i][j] != 0 {
				m[i][j] = []int{maze[i][j]}
			} else {
				m[i][j] = findPosibilities(maze, i, j)
			}
		}
	}

	// m[0] = map[int][]int{}
	// m[0][0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// m[0][1] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printPosibilities(m)

	return maze, nil
}

func makeEmptyMap() map[int]map[int][]int {
	m := map[int]map[int][]int{}
	for i := range 9 {
		m[i] = map[int][]int{}
	}
	return m
}

func printPosibilities(m map[int]map[int][]int) {
	for i := range 9 {
		for j := range 9 {
			fmt.Printf("pos [%d,%d]: %v\n", i, j, m[i][j])
		}
	}
}

func findPosibilities(maze [][]int, row, col int) []int {
	posibilities := []int{}
	existed := map[int]bool{}

	for i := range 9 {
		existed[maze[row][i]] = true
		existed[maze[i][col]] = true
	}

	startRow, startCol := (row/3)*3, (col/3)*3
	for i := range 3 {
		for j := range 3 {
			existed[maze[startRow+i][startCol+j]] = true
		}
	}

	for i := 1; i <= 9; i++ {
		if !existed[i] {
			posibilities = append(posibilities, i)
		}
	}

	return posibilities
}
