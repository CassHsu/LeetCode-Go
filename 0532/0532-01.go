func findPairs(nums []int, k int) int {
    count := 0
    m := map[int]int {}
    
    for _, n := range nums {
        if v, ok := m[n]; ok {
            m[n] = v + 1
        } else {
            m[n] = 1
        }
    }
    
    if k > 0 {
        for key, _ := range m {
            if _, ok := m[key + k]; ok {
                count++
            }
        }
    }
    
    if k == 0 {
        for _, v := range m {
            if v > 1 {
                count++
            }
        }
    }
    
    return count
}
