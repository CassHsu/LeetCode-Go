func mergeAlternately(word1 string, word2 string) string {
    ret := ""
    w1 := len(word1)
    w2 := len(word2)
    i := 0
    
    for w1 > 0 && w2 > 0 {
        ret += string(word1[i])
        ret += string(word2[i])
        w1--
        w2--
        i++
    }
    
    if w1 > 0 {
        ret += string(word1[i:])
    } else {
        ret += string(word2[i:])
    }
    
    return ret
}
