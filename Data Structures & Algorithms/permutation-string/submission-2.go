func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	var mp = make(map[byte]int)
	var mp2 = make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		mp[s1[i]]++
	}

	var left, right = 0, 0
	for right = 0; right < len(s2); right++ {
		if right < len(s1) {
			mp2[s2[right]]++
			if right == len(s1)-1 {
				if checkEq(mp, mp2) {
					return true
				}
			}
			continue
		}

		mp2[s2[left]]--
		if value, ok := mp2[s2[left]]; ok && value == 0 {
			delete(mp2, s2[left])
		}
		left++

		mp2[s2[right]]++
		if checkEq(mp, mp2) {
			return true
		}
	}

	return false
}

func checkEq(mp, mp1 map[byte]int) bool {
	if len(mp) != len(mp1) {
		return false
	}

	for key, value := range mp1 {
		if value != 0 {
			if mp[key] != value {
				return false
			}
		}
	}

	return true
}