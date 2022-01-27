import (
    "sort"
)

func countElements(nums []int) int {
    count := 0
    sort.Ints(nums)
    min, max := nums[0], nums[len(nums) - 1]
    
    for i := 1; i < len(nums); i++ {
        if max > nums[i] && min < nums[i] {
            count++
        }
    }
    
    return count
}
