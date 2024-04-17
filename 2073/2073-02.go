func timeRequiredToBuy(tickets []int, k int) int {
    seconds := 0
    n := len(tickets)
    q := make([]int, n)

    for q[k] < tickets[k] {
        for i := 0; i < n; i++ {
            if q[k] == tickets[k] && i > k {
                break
            }

            if q[i] < tickets[i] {
                q[i] += 1
                seconds += 1
            }  
        }
    }
    return seconds
}
