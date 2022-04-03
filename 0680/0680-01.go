func validPalindrome(s string) bool {
    l := 0
    r := len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return isPalindrome(s, l + 1, r) || isPalindrome(s, l, r - 1)
        }
        l++
        r--
    }
    
    return true
}

func isPalindrome(s string, i int, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    
    return true
}
