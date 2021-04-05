func generateParenthesis(n int) []string {
    ret := []string{}
    s := ""
    backtrack(n, &ret, s, 0, 0)
    return ret
}

func backtrack(n int, ret *[]string, s string, left int, right int) {
    if len(s) == n * 2 {
        *ret = append(*ret, s)
        return
    }
    
    if left > right {
        backtrack(n, ret, s + ")", left, right + 1)
    }
    
    if left < n {
        backtrack(n, ret, s + "(", left + 1, right)
    }
    return
}
