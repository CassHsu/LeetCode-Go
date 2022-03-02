func prefixCount(words []string, pref string) int {
    count := 0
    k := len(pref)
    
    for _, w := range words {
        if len(w) >= k && w[:k] == pref {
            count++
        }
    }
    
    return count
}
