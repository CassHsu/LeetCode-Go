func removeVowels(s string) string {
    vowels := map[rune]bool {'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    ans := []rune {}
    
    for _, r := range s {
        _, ok := vowels[r]
        if !ok {
            ans = append(ans, r)
        }  
    }
    
    return string(ans)
}
