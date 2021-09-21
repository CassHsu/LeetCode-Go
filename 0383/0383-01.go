func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }
    
    m := map[rune]int{}
    
    for _, c := range magazine {
        _, ok := m[c];
        if ok {
            m[c]++
        } else {
            m[c] = 1
        }
    }
    
    for _, c := range ransomNote {
        _, ok := m[c]
        if ok {
            m[c]--
            
            if m[c] < 0 {
                return false
            }
            
        } else {
            return false
        }
    }
    
    return true
}
