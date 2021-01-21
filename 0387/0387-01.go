func firstUniqChar(s string) int {
    m := map[rune]int{}
    
    for _, c := range s {
        _, ok := m[c]
        if !ok {
            m[c] = 1
        } else {
            m[c]++
        }
    }
    
    for i, c := range s {
        if m[c] == 1 {
            return i
        }
    }
    return -1
}