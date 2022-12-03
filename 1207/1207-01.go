func uniqueOccurrences(arr []int) bool {
    m := map[int]int{}
    set := make(map[int]bool)
    
    for _, a := range arr {
        if _, ok := m[a]; ok {
            m[a]++
        } else {
            m[a] = 1
        }
    }
    
    for _, v := range m {
        if _, ok := set[v]; ok {
            return false
        }
        set[v] = true
    }
    
    return true
}
