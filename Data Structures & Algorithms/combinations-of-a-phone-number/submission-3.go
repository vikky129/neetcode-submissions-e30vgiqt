func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var mp = map[string]string {
		"1":"",
		"2":"abc",
		"3":"def",
		"4":"ghi",
		"5":"jkl",
		"6":"mno",
		"7":"pqrs",
		"8":"tuv",
		"9":"wxyz",
	}

	var result = &[]string{}

	maka(digits, []rune(digits), mp, "", result)
	return *result
}

func maka(digits string, cdigits []rune, mp map[string]string, cStr string, result *[]string) {
	if len(cStr) == len(digits) {
		*result = append(*result, cStr)
		return
	}

	var cChars = mp[string(cdigits[0])]
	for i := 0; i < len(cChars); i++ {
		maka(digits, cdigits[1:], mp, cStr + string(cChars[i]), result)
	}

	return
}
