package sudoku

import (
	"errors"
	"fmt"
)

func ValidateInput(in [][]int) error {
	if len(in) != 9 {
		return errors.New("input must be 9x9")
	}

	for i := range len(in) {
		if len(in[i]) != 9 {
			return errors.New("input must be 9x9")
		}
	}

	// check for duplicates in rows
	for i := range in {
		m := map[int]bool{}
		for j := range in[i] {
			if in[i][j] == 0 {
				continue
			}
			if m[in[i][j]] {
				return fmt.Errorf("row duplicate number %d in row %d, cell %d", in[i][j], i, j)
			}
			m[in[i][j]] = true
		}
	}

	// check for duplicate in cells
	for i := range in {
		m := map[int]bool{}
		for j := range in[i] {
			if in[j][i] == 0 {
				continue
			}
			if m[in[j][i]] {
				return fmt.Errorf("cell duplicate number %d in row %d, cell %d", in[j][i], i, j)
			}
			m[in[j][i]] = true
		}
	}

	// check for duplicate in box
	// 0[0,0] 1[0,3] 2[0,6]
	// 3[3,0] 4[3,3] 5[3,6]
	// 6[6,0] 7[6,3] 8[6,6]
	for i := range in {
		a := (i / 3) * 3
		b := (i % 3) * 3
		err := validateBox(in, a, b)
		if err != nil {
			fmt.Println("err", err)
		}
	}

	return nil
}

// example start from 0, 0
// [0,0] [0,1] [0,2]
// [1,0] [1,1] [1,2]
// [2,0] [2,1] [2,2]
//
// example start from 0, 3
// [0,3] [0,4] [0,5]
// [1,3] [1,4] [1,5]
// [2,3] [2,4] [2,5]
func validateBox(in [][]int, row, cell int) error {
	m := map[int]bool{}
	for i := range 3 {
		for j := range 3 {
			if in[row+i][cell+j] == 0 {
				continue
			}
			if m[in[row+i][cell+j]] {
				return fmt.Errorf("box duplicate number %d in row %d, cell %d", in[row+i][cell+j], row+i, cell+j)
			}
			m[in[row+i][cell+j]] = true
		}
	}
	return nil
}
