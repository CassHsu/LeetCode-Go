func maxProfit(prices []int) int {
    res := 0
    size := len(prices)
    for i := 1; i < size; i++ {
        pf := prices[i] - prices[i-1]
        if pf > 0 {
            res += pf
        }
    }
    return res
}