func subsets(nums []int) [][]int {
    res := [][]int{}
    for k := 0; k <= len(nums); k++ {
        backtrack(nums, k, 0, []int{}, &res)
    }
    
    return res
}

func backtrack(nums []int, k int, first int, curr []int, res *[][]int) {
    if len(curr) == k {
        back := make([]int, len(curr))
        copy(back, curr)
        *res = append(*res, back)
    }
    
    for i := first; i < len(nums); i++ {
        curr = append(curr, nums[i])
        backtrack(nums, k, i + 1, curr, res)
        curr = curr[:len(curr) - 1]
    }
}
