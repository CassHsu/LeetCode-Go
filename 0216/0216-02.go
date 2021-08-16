func combinationSum3(k int, n int) [][]int {
    c := combination { ans: [][]int{}, k: k, n: n }
    c.backtrack([]int{}, 1, n)
    return c.ans
}

type combination struct {
    ans [][]int
    k int
    n int
}

func (c *combination) backtrack(curr []int, start int, remain int) {
    if len(curr) == c.k {
        if remain == 0 {
            tmp := make([]int, len(curr))
            copy(tmp, curr)
            c.ans = append(c.ans, tmp)
        }
        return
    }
    
    for i := start; i <= 9; i++ {
        curr = append(curr, i)
        c.backtrack(curr, i + 1, remain - i)
        curr = curr[:len(curr) - 1]
    }
}
