func splitWordsBySeparator(words []string, separator byte) []string {
    res := []string{}

    for _, word := range words {
        arr := strings.Split(word, string(separator))
        for _, w := range arr {
            if w != "" {
                res = append(res, string(w))
            }
        }
    }

    return res
}
