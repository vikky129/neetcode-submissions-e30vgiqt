func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}

	col := len(matrix[0])

	return BSearch(matrix, row, col, 0, (row*col)-1, target)
}

func BSearch(arr [][]int, row, col, start, end, value int) bool {
	mid := (start + end) / 2

	checkRow := mid / col
	checkCol := mid - (checkRow * col)
	fmt.Println(row, col, mid, checkRow, checkCol)

	if arr[checkRow][checkCol] == value {
		return true
	}

	if start == end {
		return false
	}

	if arr[checkRow][checkCol] > value {
		return BSearch(arr, row, col, start, mid, value)
	} else {
		return BSearch(arr, row, col, mid+1, end, value)
	}
}


