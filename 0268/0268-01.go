import (
    "sort"
)

func missingNumber(nums []int) int {
    sort.Ints(nums)
    size := len(nums)
    ans := 0
    
    if nums[0] != 0 {
        ans = 0
    } else if nums[size - 1] != size {
        ans = size
    } else {
        for i := 1; i < size; i++ {
            if i != nums[i] {
                ans = i
                break
            }
        }
    }
    
    return ans
}
