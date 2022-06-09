func removePalindromeSub(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    if isPalindrome(s) {
        return 1
    }
    
    return 2
}

func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    
    for l < r {
        if s[l] != s[r] {
            return false
        }
        
        l++
        r--
    }
    
    return true
}
