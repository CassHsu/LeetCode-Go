func removeVowels(s string) string {
    vowels := map[string]bool {"a": true, "e": true, "i": true, "o": true, "u": true}
    ans := ""
    
    for _, r := range s {
        str := string(r)
        _, ok := vowels[str]
        if !ok {
            ans += str
        }  
    }
    
    return ans
}
