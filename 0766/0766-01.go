func isToeplitzMatrix(matrix [][]int) bool {
    m := map[int]int{}
    
    for r, row := range matrix {
        for c, v := range row {
            val, ok := m[r - c]
            if !ok {
                m[r - c] = v
            } else if val != v {
                return false
            }
        }
    }
    
    return true
}
