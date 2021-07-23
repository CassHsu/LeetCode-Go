func findCenter(edges [][]int) int {
    center := 1
    num := len(edges)
    
    for i := 1; i <= num + 1; i++ {
        count := 0
        
        for _, e := range edges {
            if e[0] != i && e[1] != i {
                break
            }
            count++
        }
        
        if count == num {
            center = i
            break
        }
    }
    
    return center
}
