func groupAnagrams(strs []string) [][]string {

	var mp = make(map[[27]byte][]string)

	for i := 0; i< len(strs); i++ {
		var tempList = [27]byte{}
		for j := 0; j < len(strs[i]); j++ {
			tempList[strs[i][j] - 97]++
		}

		if _, ok := mp[tempList]; !ok {
			mp[tempList] = []string{}
		}

		mp[tempList] = append(mp[tempList], strs[i])
	}

    
	var fList = [][]string{}

	for _, value := range mp{
		fList = append(fList, value)
	}

	return fList
}
