func hasDuplicate(nums []int) bool {
    var mp = make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _, exists := mp[nums[i]]; exists {
            return true
        } else {
            mp[nums[i]] = 1
        }
    }

    return false
}
