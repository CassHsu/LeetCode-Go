func hIndex(citations []int) int {
    size := len(citations)
    for h := size; h > 0; h-- {
        i := size - h
        if citations[i] >= h {
            return h
        }
    }
    return 0
}
