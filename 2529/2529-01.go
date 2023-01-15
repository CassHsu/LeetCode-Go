func maximumCount(nums []int) int {
    p, n := 0, 0

    for _, num := range nums {
        if num > 0 { p++ }
        if num < 0 { n++ }
    }

    if p > n { return p }
    return n 
}
