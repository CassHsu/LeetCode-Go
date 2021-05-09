func minCostClimbingStairs(cost []int) int {
    var minCost = make([]int, len(cost) + 1)
    for i := 2; i < len(minCost); i++ {
        step1 := minCost[i-1] + cost[i-1];
        step2 := minCost[i-2] + cost[i-2];
        if step1 <= step2 {
            minCost[i] = step1
        } else {
            minCost[i] = step2
        }
    }
    return minCost[len(cost)]
}
