func findMaxK(nums []int) int {
    max := -1
    m := map[int]int{}
    
    for _, n := range nums {
        if _, ok := m[n]; ok {
            m[n] += 1
        } else {
            m[n] = 1
        }
    }
    
    for _, n := range nums {
        if n > 0 {
            if _, ok := m[-n]; ok && n > max {
                max = n
            }
        }
    }
    
    
    return max
}
