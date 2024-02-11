func findWordsContaining(words []string, x byte) []int {
    ans := []int{}
    for i, word := range words {
        find := strings.IndexByte(word, x)
        if find != -1 {
            ans = append(ans, i)
        }
    }

    return ans
}
