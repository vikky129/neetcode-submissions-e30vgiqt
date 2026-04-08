func trap(height []int) int {
	var leftMax = make([]int, len(height))
	var rightMax = make([]int, len(height))

	var maxValue = 0
	for i := 0; i < len(height); i++ {
		leftMax[i] = maxValue
		if height[i] > maxValue {
			maxValue = height[i]
		}
	}

	maxValue = 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMax[i] = maxValue
		if height[i] > maxValue {
			maxValue = height[i]
		}
	}

	var content = 0
	var totalContent = 0
	for i := 0; i < len(height); i++ {
		content = min(leftMax[i], rightMax[i]) - height[i]
		if content > 0 {
			totalContent += content
		}
	}

	return totalContent

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
