func search(nums []int, target int) int {
 return BSearch(nums, 0, len(nums) - 1, target)
}


func BSearch(nums []int,start, end, target int) int {
	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	}

	if start == end {
		return -1
	}

	if nums[mid] > target {
		return BSearch(nums, start, mid, target)
	} else {
		return BSearch(nums, mid + 1, end, target)
	}
}
