func findMin(nums []int) int {
	return BSearch(nums, 0, len(nums) - 1)
}


func BSearch(arr []int, start, end int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	for start <= end {

		if start == end {
			return arr[start]
		}
		if end == start+1 {
			if arr[end] > arr[start] {
				return arr[start]
			}

			return arr[end]
		}

		if arr[start] > arr[end] {
			start = (start + end) / 2

		} else {
			end = start
			start = 0
		}
	}

	return -1

}
