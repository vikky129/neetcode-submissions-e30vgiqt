func lengthOfLongestSubstring(s string) int {
	var mp = make(map[byte]struct{})
	var maxLen = 0
	var start, end int = 0, 0 

	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			var tStart = start
			for s[tStart] !=  s[i] {
				delete(mp, s[tStart])
				tStart++
			}
			delete(mp, s[tStart])
			if (end - start) > maxLen {
				maxLen = end - start 
			}
			start = tStart + 1
		}

		mp[s[i]] = struct{}{}
		end++
	}

	if maxLen == 0 || maxLen < len(mp){
		maxLen = len(mp)
	}

	return maxLen
}
