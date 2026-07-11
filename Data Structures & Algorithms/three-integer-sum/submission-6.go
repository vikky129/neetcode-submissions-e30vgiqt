func threeSum(nums []int) [][]int {
	var mp = make(map[int][]int)

	for i:= 0; i < len(nums); i++ {
		mp[nums[i]] = append(mp[nums[i]], i)
	}

	var currSum int
	var result = [][]int{}
	var mp1 = make(map[[3]int]struct{})
	for i := 0; i < len(nums) - 1; i++ {
		for j := i+1; j < len(nums); j++ {
			currSum = nums[i] + nums[j]
			checkSum := 0

			if currSum != 0 {
				checkSum = currSum * -1
			}
			if _, ok := mp[checkSum]; ok &&  notPresent(mp[checkSum], i,j){
				stk := st(nums[i], nums[j], currSum* -1)
				if _, ok1 := mp1[stk]; !ok1 {
					mp1[stk] = struct{}{}
					result = append(result, []int{nums[i], nums[j], checkSum})
				}
			}
		}
	}

	return result
}

func notPresent(arr []int, k , j int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != k && arr[i] != j {
			return true
		}
	}

	return false

}


func st(a, b, c int) [3]int {
	var arr = []int{a, b, c}
	sort.Ints(arr)
	var result  = [3]int{}
	result[0] = arr[0]
	result[1] = arr[1]
	result[2] = arr[2]

	return result
}


