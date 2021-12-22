func countLetters(s string) int {
    count := 1
    dp := 1
    
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            dp++
        } else {
            dp = 1
        }
        count += dp
    }
    
    return count
}
