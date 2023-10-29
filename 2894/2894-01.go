func differenceOfSums(n int, m int) int {
    s1 := 0
    s2 := 0
    for i := 1; i <= n; i++ {
        if i % m == 0 {
            s2 += i
        } else {
            s1 += i
        }
    }
    return s1 - s2
}
