# sudoku-go

Example run command
```
go run cmd/solve/main.go
```

Sample input as below, put 0 as unknown field
```
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
```

Sample output
```
6 2 7 | 8 5 1 | 3 4 9 
5 4 3 | 9 2 6 | 8 7 1 
1 9 8 | 4 7 3 | 2 5 6 
-+-+--+--+-+--+--+-+--
2 6 1 | 7 3 4 | 5 9 8 
9 7 4 | 6 8 5 | 1 3 2 
8 3 5 | 2 1 9 | 7 6 4 
-+-+--+--+-+--+--+-+--
3 1 9 | 5 6 8 | 4 2 7 
7 8 6 | 3 4 2 | 9 1 5 
4 5 2 | 1 9 7 | 6 8 3 
```
