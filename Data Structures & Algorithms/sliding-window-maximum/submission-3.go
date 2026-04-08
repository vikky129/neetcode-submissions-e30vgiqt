func maxSlidingWindow(nums []int, k int) []int {
    var max = 0
	var l, r = 0, 0
	var result = []int{}


	for r = 0; r < len(nums); r++ {
		if r < k {
			if nums[r] > max {
				max = nums[r]
			}

			if r == k - 1 {
				result = append(result, max)
			}
			continue
		}

		if nums[r] >= max {
		max = nums[r]
		result = append(result, max)
		l++
	} else if  nums[l] == max {
		l++
		var tMax = nums[l]
		for i := l; i<=r;i++ {
			if nums[i] > tMax {
				tMax= nums[i]
			}
		}
		max = tMax
		result = append(result, max)
	} else {
		result = append(result, max)
		l++
	}
	}

	

	return result
}
