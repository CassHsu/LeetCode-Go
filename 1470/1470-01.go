func shuffle(nums []int, n int) []int {
    n1 := nums[:n]
    n2 := nums[n:]
    ans := []int{}
    
    for i := range n1 {
        ans = append(ans, n1[i])
        ans = append(ans, n2[i])
    }
    
    return ans
}
