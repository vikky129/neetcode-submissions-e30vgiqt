func threeSum(nums []int) [][]int {
	QuickSort(nums, 0, len(nums) - 1)

	var mp = make(map[int]struct{})
	var res = [][]int{}

	for i := 0; i < len(nums) - 2; i++ {
		if _, ok := mp[nums[i]]; ok {
			continue
		}

		mp[nums[i]] = struct{}{}
		if nums[i] > 0 {
			break
		}

		start := i + 1
		end := len(nums) - 1
		targetSum := nums[i] * -1	

		for start < end && start >= 0 && end < len(nums) {
			if nums[start] + nums[end] < targetSum {
				start++
			} else if nums[start] + nums[end] > targetSum {
				end --
			} else {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				var nStart = start
				for nums[nStart] == nums[start] {
					nStart++
					if nStart == len(nums) - 1 {
						break
					}
				}
				start = nStart
				var nEnd = end
				for nums[nEnd] == nums[end] {
					nEnd --
					if nEnd == 0 {
						break
					}
				}

				end = nEnd
			}
		}
	}

	return res
}

func QuickSort(arr []int, start, end int) {
	if start < end {
		partition:= Partition(arr, start, end)
		QuickSort(arr, start, partition - 1)
		QuickSort(arr, partition + 1, end)
	}
}

func Partition(arr []int, start, end int) int {
	var low = start
	var pivot = arr[end]

	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr[low], arr[i] = arr[i], arr[low]
			low++
		}
	}

	arr[low], arr[end] = arr[end], arr[low]
	return low
}


