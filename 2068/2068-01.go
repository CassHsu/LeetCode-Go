func checkAlmostEquivalent(word1 string, word2 string) bool {
    c1 := map[rune]int{}
    c2 := map[rune]int{}
    s := map[rune]rune{}
    
    for _, w := range word1 { 
        s[w] = w
        v, ok := c1[w]
        if ok {
            c1[w] = v + 1
        } else {
            c1[w] = 1
        }
    }
    
    for _, w := range word2 {
        s[w] = w
        v, ok := c2[w]
        if ok {
            c2[w] = v + 1
        } else {
            c2[w] = 1
        }
    }

    for _, w := range s {
        if c1[w] - c2[w] > 3 || c2[w] - c1[w] > 3 {
            return false
        }
    }
    
    return true
}
