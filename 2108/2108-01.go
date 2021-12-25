func firstPalindrome(words []string) string {
    ans := ""
    
    for _, w := range words {
        if (isPalindrome(w)) {
            ans = w
            break
        }
    }
    
    return ans
}

func isPalindrome(w string) bool {
    i := 0
    j := len(w) - 1
    
    for i < j {
        if w[i] != w[j] {
            return false
        }
        i++
        j--
    }
    
    return true
}
