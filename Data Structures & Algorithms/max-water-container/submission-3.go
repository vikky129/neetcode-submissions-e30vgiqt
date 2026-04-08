func maxArea(heights []int) int {
	var left = 0
	var right = len(heights) - 1
	var maxH, currH = 0, 0

	for left < right {
		currH = (right - left) * (min(heights[left], heights[right]))

		if currH > maxH {
			maxH = currH
		}

		if heights[left] <= heights[right] {
			left++
		} else {
			right --
		}
	}

	return maxH
}


func min(a , b int) int {
	if a < b {
		return a
	}
	return b
}