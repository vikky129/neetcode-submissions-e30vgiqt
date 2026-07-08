func productExceptSelf(nums []int) []int {
	var firstArr = make([]int, len(nums))
	var secondArr = make([]int, len(nums))

	var i, j = 0, len(nums) - 1
	var fCurr = 1
	var lCurr = 1
	for {
		if i == len(nums) {
			break
		}
		fCurr *= nums[i]
		firstArr[i] = fCurr

		lCurr *= nums[j]
		secondArr[j]= lCurr

		i++
		j--
	}

	var finalArr = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var leftIndex = i - 1
		var rightIndex = i + 1
		var finalValue = 1
		if leftIndex >= 0 {
			finalValue = firstArr[leftIndex]
		}

		if rightIndex < len(nums) {
			finalValue *= secondArr[rightIndex]
		}

		finalArr[i] = finalValue
	}

	return finalArr
}
