func climbStairs(n int) int {
    cache := make([]int, n+1)
    return climb1or2(0, n, cache)
}

func climb1or2(i int, n int, cache []int) int {
    if i > n {
        return 0
    }
    if i == n {
        return 1
    }
    if cache[i] > 0 {
        return cache[i]
    }
    cache[i] = climb1or2(i+1, n, cache) + climb1or2(i+2, n, cache)
    return cache[i]
}
