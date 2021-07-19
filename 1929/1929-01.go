func getConcatenation(nums []int) []int {
    ans := []int{}
    
    for i := 0; i < 2; i++ {
        for _, n := range nums {
            ans = append(ans, n)
        }
    }
    
    return ans
}
