func areOccurrencesEqual(s string) bool {
    m := map[byte]int{}
    
    for i := 0; i < len(s); i++ {
        c, ok := m[s[i]]
        if ok {
            m[s[i]] = c + 1
        } else {
            m[s[i]] = 1
        }
    }
    
    num := m[s[0]]
    for _, v := range m {
        if num != v {
            return false
        }
    }
    
    return true
}
