func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    m := map[rune]int{}
    
    for _, c := range s {
        v, ok := m[c]
        if !ok {
            v = 0
        }
        m[c] = v + 1
    }
    
    for _, c := range t {
        v, ok := m[c]
        if !ok {
            return false
        }
        m[c] = v - 1
    }
    
    for _, v := range m {
        if v != 0 {
            return false
        }
    }
    
    return true
}