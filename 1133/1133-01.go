func largestUniqueNumber(nums []int) int {
    m := map[int]int{}
    singles := []int{}
    
    for _, n := range nums {
        v, ok := m[n]; if ok { 
            m[n] = v + 1 
        } else {
            m[n] = 1
        }
    }
    
    for k, v := range m {
        if v == 1 {
            singles = append(singles, k)
        }
    }
    
    if len(singles) == 0 {
        return -1
    }
    
    max := singles[0]
    for i := 1; i < len(singles); i++ {
        if singles[i] > max {
            max = singles[i]
        }
    }
    return max
}
