func findAnagrams(s string, p string) []int {
    ans := []int {}
    sn := len(s)
    pn := len(p)
    
    if sn < pn {
        return ans
    }
    
    pa := make([]int, 26)
    sa := make([]int, 26)
    
    for i := 0; i < pn; i++ {
        pa[p[i] - 'a'] += 1
    }
    
    for i := 0; i < sn; i++ {
        sa[s[i] - 'a'] += 1
        
        if i >= pn {
            sa[s[i - pn] - 'a'] -= 1
        }
        
        if isEqual(pa, sa) {
            ans = append(ans, i - pn + 1)
        }
    }
    
    return ans
}

func isEqual(pa []int, sa []int) bool {
    for i := 0; i < 26; i++ {
        if pa[i] != sa[i] {
            return false
        }
    }
    
    return true
}
