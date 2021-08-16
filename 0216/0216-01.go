func combinationSum3(k int, n int) [][]int {
    ans := [][]int{}
    backtrack(&ans, []int{}, k, n, 1, n)
    return ans
}

func backtrack(ans *[][]int, curr []int, k int, n int, start int, remain int) {
    if len(curr) == k {
        if remain == 0 {
            tmp := make([]int, len(curr))
            copy(tmp, curr)
            *ans = append(*ans, tmp)
        }
        return
    }
    
    for i := start; i <= 9; i++ {
        curr = append(curr, i);
        backtrack(ans, curr, k, n, i + 1, remain - i)
        curr = curr[:len(curr) - 1]
    }
}
