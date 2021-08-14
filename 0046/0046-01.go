func permute(nums []int) [][]int {
    ans := [][]int{}
    backtrack(&nums, []int{}, &ans)
    return ans
}

func backtrack(nums *[]int, curr []int, ans *[][]int) {
    if len(curr) == len(*nums) {
        back := make([]int, len(curr))
        copy(back, curr)
        *ans = append(*ans, back)
        
        return
    }
    
    
    for _, n := range *nums {
        if includes(curr, n) {
            continue
        }
        
        curr = append(curr, n)
        backtrack(nums, curr, ans)
        curr = curr[:len(curr) - 1]
    }
}

func includes(curr []int, n int) bool {
    for _, c := range curr {
        if c == n {
            return true
        }
    }
    return false
}
