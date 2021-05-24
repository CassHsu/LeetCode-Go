func numTrees(n int) int {
    cache := map[int]int{}
    return countTree(n, cache)
}

func countTree(n int, cache map[int]int) int {
    if n == 0 {
        return 1
    }
    if v, ok := cache[n]; ok {
        return v
    }
    
    ans := 0
    for i := 1; i <= n; i++ {
        ans += countTree(i - 1, cache) * countTree(n - i, cache)
    }
    cache[n] = ans
    return ans
}
