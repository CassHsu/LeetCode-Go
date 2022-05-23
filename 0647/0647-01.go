func countSubstrings(s string) int {
    count := 0
    size := len(s)
    
    for i := 0; i < size; i++ {
        l := i
        r := i
        for l >= 0 && r < size && s[l] == s[r] {
            count++
            l--
            r++
        }
        
        l = i
        r = i + 1
        for l >= 0 && r < size && s[l] == s[r] {
            count++
            l--
            r++
        }
    }
    
    return count
}
