func mergeAlternately(word1 string, word2 string) string {
    res := ""
    w1 := len(word1)
    w2 := len(word2)
    i := 0

    for i < w1 || i < w2 {
        if i < w1 {
            res += string(word1[i])
        }

        if i < w2 {
            res += string(word2[i])
        }

        i++
    } 

    return res
}
