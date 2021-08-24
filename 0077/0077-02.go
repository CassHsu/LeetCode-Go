func combine(n int, k int) [][]int {
    c := combType { ans: [][]int{}, n: n, k: k }
    c.backtrack([]int{}, 1,)
    return c.ans
}

type combType struct {
    ans [][]int
    k int
    n int
}

func (c *combType) backtrack(curr []int, start int) {
    if len(curr) == c.k {
        tmp := make([]int, len(curr))
        copy(tmp, curr)
        c.ans = append(c.ans, tmp)
        return
    }
    
    for i := start; i <= c.n; i++ {
        curr = append(curr, i)
        c.backtrack(curr, i + 1)
        curr = curr[:len(curr) - 1]
    }
}
