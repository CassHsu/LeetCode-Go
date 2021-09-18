import "math"

func maxProfit(prices []int) int {
    maxProfit := 0
    minCost := math.MaxInt32
    
    for i := range prices {
        if prices[i] < minCost {
            minCost = prices[i]
        }
        
        if maxProfit < (prices[i] - minCost) {
            maxProfit = prices[i] - minCost
        }
    }
    
    return maxProfit
}
