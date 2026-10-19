func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
    	rows[i] = make(map[byte]bool)
    	cols[i] = make(map[byte]bool)
    	squares[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			digit := board[r][c]
			if digit == '.' {
				continue
			}

			square := (r/3)*3 + c/3
			
			if rows[r][digit] || cols[c][digit] || squares[square][digit] {
				return false
			}

			rows[r][digit] = true
			cols[c][digit] = true
			squares[square][digit] = true

		}
	}
	return true
}
