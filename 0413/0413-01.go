func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    
    for i := 2; i < n; i++ {
        if nums[i] - nums[i - 1] == nums[i - 1] - nums[i - 2] {
            dp[i] = dp[i - 1] + 1
        }
    }
    
    sum := 0
    for _, c := range dp {
        sum += c
    }
    
    return sum
}
