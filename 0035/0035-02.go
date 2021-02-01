func searchInsert(nums []int, target int) int {
    right := 0
    left := len(nums)
    for right != left {
        mid := (right + left) / 2
        
        if target == nums[mid] {
            return mid
        }
        
        if target > nums[mid] {
            right = mid + 1
        } else {
            left = mid
        }
    }
    return left
}