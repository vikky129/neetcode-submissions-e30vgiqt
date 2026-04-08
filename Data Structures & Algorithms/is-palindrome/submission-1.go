func isPalindrome(s string) bool {

	if len(s) == 1 {
		return true
	}
	var i, j = 0, len(s) - 1

	var ty1, ty2 string
	for i <= j {

		ty1 = checkValidRange(s[i]) 
		if ty1  == "unknown"{
			i++
			continue
		}

		ty2 = checkValidRange(s[j]) 
		if ty2  == "unknown"{
			j--
			continue
		}

		if ty1 != ty2 {
			return false
		}

		if (s[i] != s[j] && !sub(s[i] - s[j])){
			return false
		}

		i++
		j--
	}

	return true

}

func sub(b byte) bool {


	if b == 32  || b == 224 {
		return true
	}


	return false
}



// 65 - 90
// 97 - 122
// 48 - 57
func checkValidRange(r byte ) string {
	if r >= 65 && r <= 90 {
		return "alpha"
	}

	if r >= 97 && r <= 122 {
		return "alpha"
	}

	if r >= 48 && r <= 57 {
		return "num"
	}

	return "unknown"
}
