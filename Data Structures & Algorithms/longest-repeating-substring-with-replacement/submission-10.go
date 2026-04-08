func characterReplacement(s string, k int) int {
	if len(s) == 0{
		return 0
	}
	var mp = make(map[byte]struct{})

	for i := 0; i < len(s); i ++ {
		mp[s[i]] = struct{}{}
	}

	var maxLen = 0
	var currLen = 0

	for key, _ := range mp {
		
		currLen = countMaxChars(key, k, s)
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen


}



func countMaxChars(ch byte, sk int, st string) int {
	var maxLen = 0
	var end = 0

	var i = 0
	var spLeft, spRight int
	var k = sk

	for i < len(st) {
		if st[i] == ch {
			spLeft = i
			spRight = i
			k = sk

			for spRight < len(st) && st[spRight] == ch {
				spRight++
				end = spRight
			}

			if i > 0 {
				for k > 0 && spLeft >= 0 {
					if st[spLeft] != ch {
						k--
					}
					spLeft--
				}

				for spLeft >= 0 && st[spLeft] == ch {
					spLeft--
				}

				if spRight < len(st) && k > 0 {
					if st[spRight] != ch {
						k--
					}
					spRight++
				}

				for spRight < len(st) && st[spRight] == ch {
					spRight++
				}

				if spLeft < 0 {
					spLeft = -1
				}

				if spRight >= len(st) {
					spRight = len(st)
				}

				if spRight-spLeft-1 > maxLen {
					fmt.Println(maxLen)
					maxLen = spRight - spLeft - 1
				}

				spLeft = i - 1
				spRight = end
				k = sk

				if spRight < len(st) && k > 0 {
					if st[spRight] != ch {
						k--
					}
					spRight++
				}

				for spRight < len(st) && st[spRight] == ch {
					spRight++
				}

				for k > 0 && spLeft >= 0 {
					if st[spLeft] != ch {
						k--
					}
					spLeft--
				}

				for spLeft >= 0 && st[spLeft] == ch {
					spLeft--
				}

				if spLeft < 0 {
					spLeft = -1
				}

				if spRight >= len(st) {
					spRight = len(st)
				}

				if spRight-spLeft-1 > maxLen {
					maxLen = spRight - spLeft - 1
					fmt.Println(maxLen)
				}

				//i = spRight

			}

		}

		i++
	}

	return maxLen

}
