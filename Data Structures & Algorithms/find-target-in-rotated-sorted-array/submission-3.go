func search(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	mI := AS(nums, 0, len(nums) - 1)

	if mI == -1 {
		return -1
	}

	if target == nums[mI] {
		return mI
	}

	if target > nums[mI] && target <= nums[len(nums) - 1]{
		return BSearch(nums, mI + 1, len(nums) - 1, target)
	}

	return BSearch(nums, 0, mI , target)
}

func BSearch(arr []int, start, end int, target int) int {
	mid := (start + end ) / 2
	if arr[mid] == target {
		return mid
	}

	if start == end {
		return -1
	}

	if end == start + 1 && arr[end] != target{
		return -1
	}

	if target < arr[mid] {
		return BSearch(arr, start, mid, target)
	} else {
		return BSearch(arr, mid + 1, end, target)
	}
}


func AS(arr []int, start, end int) int {
	for start <= end {
		if start == end {
			return start	
		}

		if end == start + 1 {
			if arr[start] < arr[end] {
				return start
			}
			return end
		}

		mid := (start + end) / 2
		if arr[start] > arr[end] {
			start = mid
		} else {
			end = start
			start = 0
		}
	}

	return -1
}
