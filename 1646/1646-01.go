func getMaximumGenerated(n int) int {
    if n == 0 { return 0 }
    
    dp := make([]int, n + 1)
    dp[1] = 1
    
    for i := 1; i < (n + 1)/2; i++ {
        dp[i * 2] = dp[i]
        dp[(i * 2) + 1] = dp[i] + dp[i + 1]
    }
    
    ans := 0
    for _, v := range dp {
        if v > ans { ans = v }
    }
    return ans
}
