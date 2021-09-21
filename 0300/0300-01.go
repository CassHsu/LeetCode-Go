func lengthOfLIS(nums []int) int {
    ans := 1
    dp := make([]int, len(nums))
    for i, _ := range dp {
        dp[i] = 1
    }
    
    for i := 0; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if dp[j] + 1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }
    
    for _, v := range dp {
        if v > ans {
            ans = v
        }
    }
    
    return ans
}
