func evalRPN(tokens []string) int {
	var mp = map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}

	var fArr = []int{}
	var value int
	var result int
	var v1, v2 int

	for i := 0; i < len(tokens); i++ {
		if _, ok := mp[tokens[i]]; ok {
			v1, v2 = fArr[len(fArr) - 2], fArr[len(fArr) - 1]
			fArr = fArr[:len(fArr) - 2]
			result = Operation(tokens[i], v1, v2)
			fArr = append(fArr, result)
		} else {
			value, _ = strconv.Atoi(tokens[i])
			fArr = append(fArr, value)
		}
	}

	return fArr[0]
}

func Operation(op string, v1, v2 int) int {
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "/":
		return int(v1 / v2)
	case "*":
		return v1 * v2
	default:
		return -1
	}
}
