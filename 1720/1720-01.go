func decode(encoded []int, first int) []int {
    ans := []int{}
    ans = append(ans, first)
    
    for _, e := range encoded {
        ans = append(ans, ans[len(ans) - 1] ^ e)
    }
    
    return ans
}