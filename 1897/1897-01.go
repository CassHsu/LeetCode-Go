func makeEqual(words []string) bool {
    n := len(words)
    m := map[rune]int{}
    
    for _, w := range words {
        for _, c := range w {
            v, ok := m[c]
            if ok {
                m[c] = v + 1
            } else {
                m[c] = 1
            }
        }   
    }
    
    for _, v := range m {
        if v % n != 0 {
            return false
        }
    }
    
    return true
}
