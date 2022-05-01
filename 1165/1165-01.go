func calculateTime(keyboard string, word string) int {
    m := map[rune]int{}
    
    for i, k := range keyboard {
        m[k] = i
    }
    
    count := 0
    curr := 0
    for _, w := range word {
        v := m[w] - curr
        if v < 0 {
            v *= -1
        }
        count += v
        curr = m[w]
    }
    
    return count
}
