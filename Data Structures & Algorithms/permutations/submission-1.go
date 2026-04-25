func permute(nums []int) [][]int {
	var result = &[][]int{}
	perm(nums, len(nums), []int{}, result)
	return *result
}

func perm(nums []int, arrLen int, cArr []int, result *[][]int) {
	if len(cArr) == arrLen {
		var temp = make([]int, arrLen)
		copy(temp, cArr)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		var arr = []int{}
		arr = append(arr, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		perm(arr, arrLen, append(cArr, nums[i]), result)
	}

	return
}
