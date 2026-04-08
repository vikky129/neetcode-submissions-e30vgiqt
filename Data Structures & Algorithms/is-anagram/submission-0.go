func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    if len(s) == 1 {
        return s[0] == t[0]
    }

   var mp = make(map[byte]int)

    for i := 0; i < len(s); i ++ {
        mp[s[i]] += 1
        mp[t[i]] -= 1
    }



    for _, value := range mp {
        if value > 0 {
            return false
        }
    }

    return true
}
