func longestConsecutive(nums []int) int {

	var storeMp = make(map[int]struct{})
	for _, value := range nums {
		storeMp[value] = struct{}{}
	}
	
	if len(storeMp) == 1 {
		return 1
	}

	var maxCount = 0
	var currCount = 1

	for key, _ := range storeMp {
		if _, ok := storeMp[key - 1]; !ok {
			currCount = MaxCount(key, storeMp, 0) + 1
			if currCount > maxCount {
				maxCount = currCount
			}
		}
	}

	return maxCount
}


func MaxCount(num int, mp map[int]struct{}, currCount int) int {
	if _, ok := mp[num + 1]; ok {
		currCount ++
		return MaxCount(num + 1, mp, currCount)
	}

	return currCount
}
