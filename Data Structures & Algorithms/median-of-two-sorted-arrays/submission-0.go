func findMedian(arr []int, isOdd bool, index int64) float64 {
	index--
	if isOdd {
		return float64(arr[index])
	}

	return (float64(arr[index]) + float64(arr[index+1])) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	index := int64(math.Ceil(float64((len(nums1) + len(nums2))) / 2))
	fmt.Println(index)
	//index
	var oddLen bool
	if (len(nums1)+len(nums2))%2 != 0 {
		oddLen = true
	}

	if len(nums1) == 0 {
		return findMedian(nums2, oddLen, index)
	}

	if len(nums2) == 0 {
		return findMedian(nums1, oddLen, index)
	}

	var i, j, cIndex int64 = 0, 0, 0
	var num int

	for cIndex < index {

		if i < int64(len(nums1)) && nums1[i] <= nums2[j] {
			num = nums1[i]
			i++
		} else {
			num = nums2[j]
			j++
		}
		cIndex++

		if cIndex == index {
			if oddLen {
				return float64(num)
			}
			var num2 int
			if i < int64(len(nums1)) && nums1[i] <= nums2[j] {
				num2 = nums1[i]
				i++
			} else {
				num2 = nums2[j]
				j++
			}

			return float64((float64(num) + float64(num2)) / 2)

		}
	}

	return -1

}
