func findUnsortedSubarray(nums []int) int {
    r := 0
    l := len(nums)
    
    for i := 0; i < len(nums) - 1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[j] < nums[i] {
                if j > r {
                    r = j
                }
                if i < l {
                    l = i
                }
            }
        }
    }
    
    if r - l > 0 {
        return r - l + 1
    } else { return 0 }
}
