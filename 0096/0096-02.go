func numTrees(n int) int {
    cache := make([]int, n + 1)
    cache[0] = 1
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            cache[i] += cache[j - 1] * cache[i - j]
        }
    }
    return cache[n]
}
