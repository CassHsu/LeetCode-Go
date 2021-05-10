var cache map[int]int

func minCostClimbingStairs(cost []int) int {
    cache = make(map[int]int)
    return minCost(cost, len(cost))
}

func minCost(cost []int, i int) int {
    if i <= 1 {
        return 0
    }
    v, ok := cache[i]
    if ok {
        return v
    }
    
    step1 := cost[i-1] + minCost(cost, i-1)
    step2 := cost[i-2] + minCost(cost, i-2)
    
    if step1 <= step2 {
        cache[i] = step1
    } else {
        cache[i] = step2
    }
    return cache[i]
}
