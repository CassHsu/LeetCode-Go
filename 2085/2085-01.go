func countWords(words1 []string, words2 []string) int {
    count := 0
    m := map[string][]int{}
    
    for _, w := range words1 {
        if v, ok := m[w]; ok {
            v[0]++
            m[w] = v
        } else {
            m[w] = []int{1, 0}
        }
    }
    
    for _, w := range words2 {
        if v, ok := m[w]; ok {
            v[1]++
            m[w] = v
        } else {
            m[w] = []int{0, 1}
        }
    }
    
    for _, v := range m {
        if v[0] == 1 && v[1] == 1 {
            count++
        }
    }
    
    return count
}
