package main

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool

	for r := range board {
		for c := range board[0] {
			cell := board[r][c]
			if cell != '.' {
				b := (r/3)*3 + (c / 3)
				idx := int(cell - '1')
				if rows[r][idx] || cols[c][idx] || boxes[b][idx] {
					return false
				}
				rows[r][idx] = true
				cols[c][idx] = true
				boxes[b][idx] = true
			}
		}
	}
	return true
}
