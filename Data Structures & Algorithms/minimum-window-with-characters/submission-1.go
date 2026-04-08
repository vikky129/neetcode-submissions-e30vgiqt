func minWindow(s string, t string) string {
	if len(t) == 0 && len(s) != 0 {
		return ""
	}

	var fStr = ""

	if len(s) == 0 && len(t) > 0 {
		return ""
	}

	if len(s) == 0 && len(t) == 0 {
		return ""
	}

	var mpt = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mpt[t[i]]++
	}


	var mps = make(map[byte]int)
	var l, r = 0, 0
	for r = 0; r < len(s); r++ {
		if r < len(t) {
			mps[s[r]]++

			if r == len(t)-1 {
				if compare(mpt, mps) {
					if r+1-l < len(fStr) || fStr == "" {
						fStr = s[l : r+1]
					}
					mps[s[l]]--
					if mps[s[l]] == 0 {
						delete(mps, s[l])
					}
					l++

				}
			}
			continue
		}

		for mpt[s[l]] == 0 || mps[s[l]] > mpt[s[l]] {
			mps[s[l]]--
			if mps[s[l]] == 0 {
				delete(mps, s[l])
			}
			l++

			if l >= len(s) {
				l = len(s) - 1
				break
			}
		}

		mps[s[r]]++
		if compare(mpt, mps) {
			if r+1-l < len(fStr) || fStr == "" {
				fStr = s[l : r+1]
			}
			mps[s[l]]--
			if mps[s[l]] == 0 {
				delete(mps, s[l])
			}
			l++
		}

	}

	return fStr
}

func compare(mpt, mps map[byte]int) bool {

	for key, value := range mpt {
		if mps[key] < value {
			return false
		}
	}

	return true
}