func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProf := 0
    
    for _, p := range prices {
        if p < minPrice {
            minPrice = p
        }
        
        if maxProf < p - minPrice {
            maxProf = p - minPrice
        }
    }
    
    return maxProf
}
