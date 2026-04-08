func minEatingSpeed(piles []int, h int) int {
	max, sum := maxElement(piles)

	if len(piles) == h {
		return max
	}

	var pv = sum/h

	if pv == 0 {
		pv = 1
	}

	var check = false

	for !check {
		check = checkPossibility(piles, pv, h)
		pv++
	}

	return pv - 1
}

func checkPossibility(arr []int, pv, h int) bool {
	var ch = 0
	var div int
	for i := 0; i < len(arr); i++ {
		if arr[i] <= pv {
			ch++
		} else {
			div = arr[i] / pv
			ch += div
		}

		if div > 0 && (arr[i]-(div*pv)) > 0 {
			ch++
		}
	}

	if ch <= h {
		return true
	}

	return false
}

func maxElement(arr []int ) (int, int) {
	var sum = arr[0]
	var max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		sum += arr[i]
	}

	return max, sum
}
