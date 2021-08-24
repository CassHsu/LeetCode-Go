func combine(n int, k int) [][]int {
    ans := [][]int{}
    backtrack(&ans, []int{}, 1, n, k)
    return ans
}

func backtrack(ans *[][]int, curr []int, start int, n int, k int) {
    if len(curr) == k {
        tmp := make([]int, len(curr))
        copy(tmp, curr)
        *ans = append(*ans, tmp)
				return
    }
    
    for i := start; i <= n; i++ {
        curr = append(curr, i)
        backtrack(ans, curr, i + 1, n, k)
        curr = curr[:len(curr) - 1]
    }
}
