func generateParenthesis(n int) []string {
	var result = &[]string{}
	var cArr = []string{}

	parenthesis(n, n, cArr, result)
	return *result
}

func parenthesis(oCount, eCount int, cArr []string, result *[]string)  {
	if eCount == 0 && oCount == 0 {
		*result = append(*result, strings.Join(cArr, ""))
		return
	}

	if len(cArr) == 0 {
		cArr = append(cArr, "(")
		parenthesis(oCount - 1, eCount, cArr, result)
		return
	}

	if eCount == 1 && oCount == 0 {
		cArr = append(cArr, ")")
		parenthesis(oCount, eCount - 1, cArr, result)
		return
	}

	if oCount > 0 {
		parenthesis(oCount - 1, eCount, append(cArr, "("), result)
	}

	if eCount > oCount {
		parenthesis(oCount, eCount - 1, append(cArr, ")"), result)
	}

	return 
}
