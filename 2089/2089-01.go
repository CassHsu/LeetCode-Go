import (
    "sort"
)

func targetIndices(nums []int, target int) []int {
    ans := []int{}
    sort.Ints(nums)
    
    for i, n := range nums {
        if n == target {
            ans = append(ans, i)
        }
    }
    
    return ans
}
