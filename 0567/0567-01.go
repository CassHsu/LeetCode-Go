func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    countS1 := [26]int{}
    for i := 0; i < len(s1); i++ {
        countS1[s1[i] - 'a']++
    }
    
    for i := 0; i <= len(s2) - len(s1); i++ {
        countS2 := [26]int{}
        for j := 0; j < len(s1); j++ {
            countS2[s2[i + j] - 'a']++    
        }
        
        if match(countS1, countS2) {
            return true
        }
    }
    return false
}

func match(c1 [26]int, c2 [26]int) bool {
    for i := 0; i < 26; i++ {
        if c1[i] != c2[i] {
            return false
        }
    }
    return true
}
