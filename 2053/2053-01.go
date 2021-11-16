func kthDistinct(arr []string, k int) string {
    count := map[string]int {}
    
    for _, a := range arr {
        if _, ok := count[a]; ok {
            count[a]++
        } else {
            count[a] = 1
        }
    }
    
    for _, a := range arr {
        if count[a] == 1 { 
            k--
            if k == 0 {
                return a
            }
        }
    }
    
    return ""
}
