func sortArrayByParity(nums []int) []int {
    l, r := 0, len(nums) - 1
    
    for l < r {
        if (nums[l] % 2 > nums[r] % 2) {
            nums[l], nums[r] = nums[r], nums[l]
        }
        
        if nums[l] % 2 == 0 {
            l++
        }
        
        if nums[r] % 2 == 1 {
            r--
        }
    }
    
    return nums
}
