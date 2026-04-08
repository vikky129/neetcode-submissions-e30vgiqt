func isValid(s string) bool {

	if len(s) == 1 {
		return false
	}

    var mp = map[byte]byte{
		41:  40,
		93:  91,
		125: 123,
	}

	var newArr = []byte{}

	for i := 0; i < len(s); i++ {
		if value, ok := mp[s[i]]; ok {
			if len(newArr) == 0 {
				return false
			}
			if newArr[len(newArr)-1] != value {
				return false
			} else {
				newArr = newArr[:len(newArr)-1]
			}
		} else {
			newArr = append(newArr, s[i])
		}
	}

	if len(newArr) == 0 {
		return true
	} else {
		return false
	}
}
