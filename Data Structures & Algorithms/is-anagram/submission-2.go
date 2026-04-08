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

        if mp[s[i]] == 0 {
            delete(mp, s[i])
        }

        if mp[t[i]] == 0 {
            delete(mp, t[i])
        }
    }

    return len(mp) ==0
}
