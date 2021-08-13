func subsets(nums []int) [][]int {
    res := [][]int{}
    backtrack(nums, 0, []int{}, &res)
    return res
}

func backtrack(nums []int, i int, curr []int, res *[][]int) {
    if i >= len(nums) {
        back := make([]int, len(curr))
        copy(back, curr)
        *res = append(*res, back)
        
        return
    }
    
    curr = append(curr, nums[i])
    backtrack(nums, i + 1, curr, res)
    
    curr = curr[:len(curr)-1]
    backtrack(nums, i + 1, curr, res)
}
