func topKFrequent(nums []int, k int) []int {

	return bucketSort(nums, k)

	// if len(nums) == k {
	// 	return nums
	// }

	// var arr = []int{}
	// var mp = make(map[int]int)

	// for i := 0; i < len(nums); i++ {
	// 	mp[nums[i]] ++
	// }





	// for i := 0; i < k; i++ {
	// 	arr = append(arr, returnMax(mp))
	// }

	// return arr
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

// Implementation with BucketSort

func bucketSort(arr []int, k int) []int {
	var mp = make(map[int]int)

	for i := 0; i < len(arr); i++ {
		mp[arr[i]]++
	}

	var freq = make([][]int, len(arr) + 1)
	for num, count := range mp {
		freq[count] = append(freq[count], num)
	}

	var res = []int{}
	for i := len(freq) - 1; i >= 0; i-- {
		for j := 0; j < len(freq[i]); j++ {
			res = append(res, freq[i][j])
			if len(res) == k {
				return res
			}

		}

	}
return res
}


