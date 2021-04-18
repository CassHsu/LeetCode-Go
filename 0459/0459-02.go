func repeatedSubstringPattern(s string) bool {
    size := len(s)
    for i := size / 2; i > 0; i-- {
        if size % i == 0 {
            repeat := size / i
            substr := ""
            for j := 0; j < repeat; j++ {
                substr += s[:i]
            }
            
            if substr == s {
                return true
            }
        }
    }
    return false
}
