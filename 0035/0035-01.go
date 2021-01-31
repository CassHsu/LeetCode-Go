func searchInsert(nums []int, target int) int {
    size := len(nums)
    if target > nums[size - 1] {
        return size
    }
    
    res := 0
    for i, n := range nums {
        if n >= target {
            res = i
            break
        }
    }
    return res
}