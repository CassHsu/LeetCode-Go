import (
    "sort"
)

func findUnsortedSubarray(nums []int) int {
    r := 0
    l := len(nums)
    
    snums := make([]int, len(nums))
    copy(snums, nums)
    sort.Ints(snums)
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != snums[i] {
            if i > r {
                r = i
            }
            if i < l {
                l = i
            }
        }
    }
    
    if r - l > 0 {
        return r - l + 1
    } else { return 0 }
}
