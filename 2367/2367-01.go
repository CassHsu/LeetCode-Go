func arithmeticTriplets(nums []int, diff int) int {
    count := 0
    m := map[int]bool{}

    for _, n := range nums {
        if _, ok1 := m[n - diff]; ok1 {
            if _, ok2 := m[n - diff -diff]; ok2 {
                count++
            }
        }
        m[n] = true
    }

    return count
}
