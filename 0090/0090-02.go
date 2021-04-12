import (
    "sort"
)

func subsetsWithDup(nums []int) [][]int {
    ret := [][]int{[]int{}}
    var curr []int
    
    sort.Ints(nums)
    backtrack(nums, curr, &ret)
    return ret
}

func backtrack(nums []int, curr []int, ret *[][]int) {
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        curr = append(curr, nums[i])
        ss := make([]int, len(curr))
        copy(ss, curr)
        
        *ret = append(*ret, ss)
        
        backtrack(nums[i+1:], curr, ret)
        curr = curr[:len(curr) - 1]
    }
}
