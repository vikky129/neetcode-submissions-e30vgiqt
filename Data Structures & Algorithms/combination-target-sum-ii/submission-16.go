 import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	var mp = map[string][]int{}
	slices.Sort(candidates)
	mp = comb(candidates, 0, 0, target, []int{}, mp)
	var result = [][]int{}
	for _, val := range mp {
		result = append(result, val)
	}

	return result
}

func comb(nums []int, curr int, index int, target int, cArr []int, mp map[string][]int) map[string][]int {
	if curr > target {
		return mp
	}

	if curr == target {
		temp := make([]int, len(cArr))
        copy(temp, cArr)
		key := ""
		//sort.Ints(temp)
		for k := 0; k < len(temp); k++ {
			key += fmt.Sprintf("%d", temp[k])
		}
		mp[key] = temp
	}

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
      		continue
		}
		if curr + nums[i] > target {
			break
		}
		mp = comb(nums, curr+nums[i], i+1, target, append(cArr, nums[i]), mp)
	}

	return mp
}
