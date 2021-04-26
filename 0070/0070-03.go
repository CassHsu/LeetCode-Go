func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    
    prev := 1
    curr := 2
    for i := 3; i <= n; i++ {
        prev, curr = curr, prev + curr
    }
    return curr
}
