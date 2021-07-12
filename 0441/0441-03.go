func arrangeCoins(n int) int {
    l := 0
    r := n
    
    for l <= r {
        m := l + (r - l) / 2
        curr := m * (m + 1) / 2
        
        if (curr == n) {
            return m
        }
        
        if (n < curr) {
            r = m - 1
        } else {
            l = m + 1
        }
    }
    
    return r
}
