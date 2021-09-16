func fib(n int) int {
    if n < 2 { return n }
    
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    
    for i := 2; i < len(dp); i++ {
        dp[i] = -1
    }
    
    return fibHelper(n, dp)
}

func fibHelper(n int, dp []int) int {
    if dp[n] != -1 { return dp[n] }
    
    dp[n] = fibHelper(n - 1, dp) + fibHelper(n - 2, dp)
    return dp[n]
}
