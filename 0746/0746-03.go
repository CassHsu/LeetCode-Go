func minCostClimbingStairs(cost []int) int {
    s1 := 0
    s2 := 0
    
    for i := 2; i <= len(cost); i++ {
        tmp := s1
        c1 := s1 + cost[i-1]
        c2 := s2 + cost[i-2]
        if c1 <= c2 {
            s1 = c1
        } else {
            s1 = c2
        }
        s2 = tmp
        
    }
    return s1
}
