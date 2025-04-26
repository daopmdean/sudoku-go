package sudoku

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	in := [][]int{
		{6, 0, 0, 8, 0, 0, 3, 0, 0},
		{5, 0, 0, 0, 2, 0, 0, 0, 0},
		{1, 0, 0, 4, 0, 3, 0, 5, 6},
		{0, 6, 1, 0, 0, 0, 5, 9, 8},
		{9, 0, 4, 6, 0, 5, 0, 0, 2},
		{8, 3, 5, 0, 0, 9, 7, 0, 0},
		{3, 0, 9, 5, 0, 8, 4, 0, 0},
		{7, 8, 0, 0, 4, 2, 9, 1, 0},
		{0, 0, 2, 0, 0, 0, 0, 8, 0},
	}

	_, err := Solve(in)
	if err != nil {
		t.Error(err)
	}
}

func TestSolve2(t *testing.T) {
	in := [][]int{
		{8, 7, 0, 0, 0, 0, 6, 0, 1},
		{0, 9, 0, 0, 0, 0, 2, 7, 5},
		{0, 1, 0, 0, 6, 2, 0, 4, 8},
		{0, 0, 6, 1, 0, 0, 0, 0, 7},
		{1, 5, 0, 0, 0, 0, 0, 0, 2},
		{3, 0, 0, 0, 0, 8, 5, 0, 6},
		{4, 6, 9, 0, 0, 0, 8, 0, 0},
		{2, 8, 0, 9, 3, 0, 0, 5, 0},
		{0, 3, 0, 0, 2, 4, 1, 6, 9},
	}
	_, err := Solve(in)

	print(in)
	if err != nil {
		t.Error(err)
	}
}

func print(maze [][]int) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(maze[i][j], " ")
			if j == 2 || j == 5 {
				fmt.Print("| ")
			}
		}
		if i == 2 || i == 5 {
			fmt.Print("\n-+-+--+--+-+--+--+-+--\n")
		} else {
			fmt.Print("\n")
		}
	}
}

func empty() [][]int {
	sl := make([][]int, 9)
	for i := range 9 {
		sl[i] = make([]int, 9)
	}
	return sl
}
