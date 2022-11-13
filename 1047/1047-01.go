func removeDuplicates(s string) string {
    res := []rune{}

    for _, r := range s {
        size := len(res)

        if size > 0 && res[size - 1] == r {
            res = res[:size - 1]
        } else {
            res = append(res, r)
        }
    }

    return string(res)
}
