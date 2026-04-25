import "slices"

func subsetsWithDup(nums []int) [][]int {
	var result = &[][]int{}
	tset := [][]int{}
	tset = append(tset, []int{})
	mp := make(map[string]struct{})
	set(nums, 0, tset,mp, result)
	return *result
}

func set(nums []int, idx int, tset [][]int,mp map[string]struct{}, result *[][]int) {
	if idx >= len(nums) {
		return
	}
	l := len(tset)

	for i := 0; i < l; i++ {
		karr := make([]int, len(tset[i]))
		copy(karr, tset[i])

		karr = append(karr, nums[idx])
		slices.Sort(karr)
		st := ""
		for j := 0; j < len(karr); j++ {
			st += fmt.Sprintf("%d", karr[j])
		}

		if _, ok := mp[st]; ok {
			continue
		}
		mp[st] = struct{}{}
		tset = append(tset, karr)
	}
	*result = tset
	set(nums, idx + 1, tset, mp, result)
}
