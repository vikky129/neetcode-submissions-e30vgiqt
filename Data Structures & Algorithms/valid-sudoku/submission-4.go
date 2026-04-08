func isValidSudoku(board [][]byte) bool {
	var isValid bool

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j += 3 {
			isValid = checkBox(board, i, j)
			if !isValid {
				return false
			}
		}
	}
	return true
}

func checkBox(board [][]byte, row, col int) bool {
	var mp = make(map[byte]bool)

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if _, ok := mp[board[i][j]]; ok {
				fmt.Println("1")
				return false

			} else {
				if board[i][j] != 46{
					mp[board[i][j]] = true
				}
			}

			var isValid = checkRow(board, board[i][j], j, i)
			if !isValid {
				return false
			}

			isValid = checkCol(board, board[i][j], i, j)
			if !isValid {
				return false
			}
		}
	}

	return true
}

func checkRow(board [][]byte, value byte, col int, index int) bool {
	if value == 46 {
		return true
	}
	for i := 0; i < 9; i++ {
		if i == index {
			continue
		}
		if board[i][col] == value {
			fmt.Println("2", col, index, value)
			return false
		}
	}

	return true
}

func checkCol(board [][]byte, value byte, row int, index int) bool {
	if value == 46 {
		return true
	}
	for i := 0; i < 9; i++ {
		if i == index {
			continue
		}
		if board[row][i] == value {
			fmt.Println("3")
			return false
		}
	}

	return true
}