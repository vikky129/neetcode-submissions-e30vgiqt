// 1,2,3,2,2

func findDuplicate(nums []int) int {
	var index int
    for i := 0; i < len(nums); i++ {

		index = nums[i]
		// 4
		if index < 0 {
			index = index * -1
		}
		if nums[index] < 0 {
			fmt.Println(i)
			return index
		}

		nums[index] = -1 * nums[index]
	}

	return -1
}
