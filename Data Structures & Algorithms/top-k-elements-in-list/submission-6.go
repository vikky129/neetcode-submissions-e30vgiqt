func topKFrequent(nums []int, k int) []int {

	if len(nums) == k {
		return nums
	}

	var arr = []int{}
	var mp = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] ++
	}





	for i := 0; i < k; i++ {
		arr = append(arr, returnMax(mp))
	}

	return arr
}


func returnMax(mp map[int]int ) int {
	var max = 0
	var aValue = 0

	for key, value := range mp {
		if value > max {
			max = value
			aValue = key	
		}
	}

	delete(mp, aValue)

	return aValue
}


