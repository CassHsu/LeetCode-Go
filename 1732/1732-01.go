func largestAltitude(gain []int) int {
    max := 0
    last := 0
    for _, g := range gain {
        last = last + g
        if last > max {
            max = last
        }
    }
    return max
}
