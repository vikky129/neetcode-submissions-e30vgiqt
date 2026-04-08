func twoSum(nums []int, target int) []int {
    var mp = make(map[int][]int)

    for i := 0; i < len(nums); i++ {
        if _, ok := mp[i]; ok {
            mp[nums[i]] = append(mp[nums[i]], i)
        } else {
            mp[nums[i]] = []int {i}
        }
    }

    for i := 0; i < len(nums); i++ {
        if value, ok := mp[target - nums[i]]; ok {
            if len(value) ==1 && value[0] != i {
                return []int{i, value[0]}
            }

            for j := 0; j < len(value); j++ {
                if value[j] != i {
                    return []int{i, value[j]}
                }
            }
        }
    }

    return []int{}
}
