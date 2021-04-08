func findDuplicate(nums []int) int {
    slow := 0
    fast := 0
    for true {
        fast = nums[fast]
        fast = nums[fast]
        
        slow = nums[slow]
        
        if slow == fast {
            slow = 0
            break
        }
    }
    
    for true {
        fast = nums[fast]
        slow = nums[slow]
        
        if slow == fast {
            return slow
        }
    }
    return -1
}
