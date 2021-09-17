func maxSubArray(nums []int) int {
    if len(nums) == 0 { return 0 }
    
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    ans := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] + dp[i - 1] >= nums[i] {
            dp[i] = nums[i] + dp[i - 1]
        } else {
            dp[i] = nums[i]
        }
        
        if dp[i] > ans { ans = dp[i] }
    }
    
    return ans
}
