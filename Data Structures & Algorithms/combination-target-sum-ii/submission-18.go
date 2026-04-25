import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var result [][]int
	comb(candidates, 0, 0, target, []int{}, &result)

	return result
}

func comb(nums []int, curr int, index int, target int, cArr []int, result *[][]int) {
	if curr > target {
		return
	}

	if curr == target {
		temp := make([]int, len(cArr))
		copy(temp, cArr)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(nums); i++ {

		// ✅ skip duplicates at same level
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		if curr+nums[i] > target {
			break
		}

		comb(nums, curr+nums[i], i+1, target, append(cArr, nums[i]), result)
	}
}