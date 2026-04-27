func partition(s string) [][]string {
	result := &[][]string{}

	pn(s, 0, []string{}, result)
	return *result
}

func pn(s string, idx int, cArr []string, result *[][]string) {
	if idx >= len(s) {
		temp := make([]string, len(cArr))
		copy(temp, cArr)
		*result = append(*result, temp)
		return
	}

	for j := idx + 1; j <= len(s); j++ {
		if checkPalindrome(s[idx : j]) {
			pn(s, j, append(cArr, s[idx : j]), result)
		}
	}

	return
}


func checkPalindrome(st string) bool {
	if len(st) == 1 {
		return true
	}

	var s, e = 0, len(st) - 1
	for st[s] == st[e] {
		if s == e || e == s + 1 {
			return true
		}

		s++
		e--
	}

	return false
}