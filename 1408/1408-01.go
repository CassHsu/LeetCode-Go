func stringMatching(words []string) []string {
    ans := make([]string, 0)
    
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i == j {
                continue
            }
            
            if strings.Contains(words[j], words[i]) {
                ans = append(ans, words[i])
                break
            }
        }
    }
    
    return ans
}
