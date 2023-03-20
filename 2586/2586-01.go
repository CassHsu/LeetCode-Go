func vowelStrings(words []string, left int, right int) int {
    count := 0
    vowels := map[string]bool{
        "a": true,
        "e": true,
        "i": true,
        "o": true,
        "u": true}

    for i := left; i <= right; i++ {
        w := words[i]
        end := len(w) - 1
        if vowels[string(w[0])] && vowels[string(w[end])] {
            count++
        }
    }

    return count
}
