func buildArray(nums []int) []int {
    ans := []int{}
    
    for _, n := range nums {
        ans = append(ans, nums[n])
    }
    
    return ans
}
