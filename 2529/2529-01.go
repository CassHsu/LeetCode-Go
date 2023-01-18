func maximumCount(nums []int) int {
    p, n := 0, 0

    for _, num := range nums {
        if num > 0 { 
            p++
            continue
        }
        if num < 0 { 
            n++
            continue
        }
    }

    if p > n { return p }
    return n 
}
