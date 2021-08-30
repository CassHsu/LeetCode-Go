func maxAscendingSum(nums []int) int {
    count := nums[0]
    ans := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            count += nums[i]
        } else {
            count = nums[i]
        }
        
        if count > ans {
            ans = count
        }
    }
    
    return ans
}
