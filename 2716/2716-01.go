func minimizedStringLength(s string) int {
    set := make(map[rune]bool)

    for _, c := range s {
        _, exists := set[c]
        if !exists {
            set[c] = true
        }
    }

    return len(set)
}
