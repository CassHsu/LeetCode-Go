func timeRequiredToBuy(tickets []int, k int) int {
    ans := 0
    count := tickets[k]
    n := len(tickets)
    
    for r := 0; r < count; r++ {
        for c := 0; c < n; c++ {
            if r == count - 1 && c > k {
                break
            }
            
            if tickets[c] > 0 {
                ans++
                tickets[c]--
            }
        }
    }
    
    return ans
}
