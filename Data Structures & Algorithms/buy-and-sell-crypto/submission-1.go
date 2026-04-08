func maxProfit(prices []int) int {
	var minArr = make([]int, len(prices))
	var maxArr = make([]int, len(prices))

	var minValue = prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		minArr[i] = minValue
	}

	var maxValue = prices[len(prices) - 1]
	for j := len(prices) - 1; j >= 0; j-- {
		if prices[j] > maxValue {
			maxValue = prices[j]
		}

		maxArr[j] = maxValue
	}

	maxValue = 0
	var currValue int
	for i := 0; i < len(prices) - 1; i++ {
		currValue = maxArr[i + 1] - minArr[i]

		if currValue > maxValue {
			maxValue = currValue
		}
	}


	return maxValue
}
