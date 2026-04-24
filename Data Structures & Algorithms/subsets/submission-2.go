func subsets(nums []int) [][]int {
	var result = [][]int{}

	result = append(result, []int{})

	var newSub = [][]int{}

	for i := 0; i < len(nums); i++ {
		newSub = [][]int{}
		for j := 0; j < len(result); j++ {
			sub := make([]int, len(result[j]))
			copy(sub, result[j])
			newSub = append(newSub, append(sub, nums[i]))
		}

		result = append(result, newSub...)
	}

	return result
}




