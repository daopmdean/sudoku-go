package sudoku

// 5 3 _ | _ 7 _ | _ _ _
// 6 _ _ | 1 9 5 | _ _ _
// _ 9 8 | _ _ _ | _ 6 _
// ------+-------+------
// 8 _ _ | _ 6 _ | _ _ 3
// 4 _ _ | 8 _ 3 | _ _ 1
// 7 _ _ | _ 2 _ | _ _ 6
// ------+-------+------
// _ 6 _ | _ _ _ | 2 8 _
// _ _ _ | 4 1 9 | _ _ 5
// _ _ _ | _ 8 _ | _ 7 9

func Solve(in [][]int) ([][]int, error) {
	if err := ValidateInput(in); err != nil {
		return nil, err
	}

	// m := map[int]map[int][]int{}
	// m[0] = map[int][]int{}
	// m[0][0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// m[0][1] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(m)

	if solveBacktracking(in) {
		return in, nil
	}

	return in, nil
}
