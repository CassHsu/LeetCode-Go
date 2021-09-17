func maxSubArray(nums []int) int {
    if len(nums) == 0 { return 0 }
    
    dp_0 := nums[0]
    dp_1 := 0
    ans  := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] + dp_0 >= nums[i] {
            dp_1 = nums[i] + dp_0
        } else {
            dp_1 = nums[i]
        }
        dp_0 = dp_1
        
        if dp_1 > ans { ans = dp_1 }
    }
    
    return ans
}
