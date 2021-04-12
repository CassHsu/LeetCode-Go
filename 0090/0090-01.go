import (
    "sort"
)

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    
    ret := [][]int{};
    ret = append(ret, []int{})
    start, end := 0, 0
    
    for i, n := range nums {
        start = 0
        if i > 0 && nums[i-1] == nums[i] {
            start = end + 1
        }
        end = len(ret) - 1
        
        for j := start; j <= end; j++ {
            ss := []int{}
            ss = append(ss, ret[j]...)
            ss = append(ss, n)
            ret = append(ret, ss)
        }
    }
    return ret
}
