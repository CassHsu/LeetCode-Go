//Bottom Up
func coinChange(coins []int, amount int) int {
    if amount == 0 || len(coins) == 0 {
        return 0
    }
    
    add1 := amount + 1
    cache := []int{}
    
    cache = append(cache, 0)
    for i := 1; i <= amount; i++ {
        cache = append(cache, add1)
    }
    
    for _, c := range coins {
        for i := c; i <= amount; i++ {
            if cache[i - c] + 1 < cache[i] {
                cache[i] = cache[i - c] + 1
            }
        }
    }
    
    if cache[amount] >= add1 {
        cache[amount] = -1
    }
    return cache[amount]
}