func countWords(words1 []string, words2 []string) int {
    count := 0
    
    m := make(map[string]int)
    for _, w := range words1 {
        m[w]++
    }
    
    for _, w := range words2 {
        if v := m[w]; v < 2 {
            m[w]--
        }
    }
    
    for _, v := range m {
        if v == 0 {
            count++
        }
    }
    
    return count
}
