func isMonotonic(nums []int) bool {
    mode := -1
    
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] < nums[i] {
            if mode == 1 {
                return false
            } else {
                mode = 0
            }
        }
        
        if nums[i - 1] > nums[i] {
            if mode == 0 {
                return false
            } else {
                mode = 1
            }
        }
    }
    
    return true
}
