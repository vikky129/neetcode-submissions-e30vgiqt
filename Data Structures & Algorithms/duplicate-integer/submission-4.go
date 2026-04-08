func hasDuplicate(nums []int) bool {
    var mp = make(map[int]struct{}, len(nums))
    return quicksortInPlace(nums, 0, len(nums) - 1, mp)
}

func quicksortInPlace(arr []int, low, high int, mp map[int]struct{}) bool {
	if low < high {
		pivotIdx, found := partition1(arr, low, high, mp)
		if found {
			return true
		}
		quicksortInPlace(arr, low, pivotIdx-1, mp)
		quicksortInPlace(arr, pivotIdx+1, high, mp)
	}

	return false
}

func partition1(arr []int, low, high int, mp map[int]struct{}) (int, bool) {
	var pivot = arr[high]
	pivotPosition := high

    if _, ok := mp[pivot]; ok {
		return -1, true
	} else {
		mp[pivot] = struct{}{}
	}

	for i := low; i < high; i++ {
		if i >= pivotPosition {
			break
		}

		if _, ok := mp[arr[i]]; ok {
			return -1, true
		} else {
			mp[arr[i]] = struct{}{}
		}
		if arr[i] > pivot {
			arr[i], arr[pivotPosition] = arr[pivotPosition], arr[i]
			pivotPosition--
		}
	}

	return pivotPosition, false
}

