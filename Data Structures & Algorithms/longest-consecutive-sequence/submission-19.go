func longestConsecutive(nums []int) int {

	return FindStartElement(nums)
}


func MaxCount(num int, mp map[int]struct{}, currCount int) int {
	if _, ok := mp[num + 1]; ok {
		currCount ++
		return MaxCount(num + 1, mp, currCount)
	}

	return currCount
}

func FindStartElement(arr []int) int {
	var starter = make(map[int]*int)
	var mp = make(map[int]struct{})
	var result []int

	for i:= 0; i < len(arr); i++ {
		mp[arr[i]] = struct{}{}
	}

	// 1,7,6,4,3,2
	for i := 0; i < len(arr); i++ {
		starter[arr[i]] = nil
		if _, ok := mp[arr[i] - 1]; ok {
			starter[arr[i]] = new(int)
		}
	}

	for key, value := range starter {
		if value == nil {
			result = append(result, key)
		}
	}

	var count = 0

	fmt.Println(result)

	for _, start := range result {
		newCount := 1
		var newStart = start
		for {
			if _, ok := mp[newStart+ 1]; !ok {
				break
			}
			newStart ++
			newCount++
		}

		if newCount > count {
			count = newCount
		}
	}

	return count

}
