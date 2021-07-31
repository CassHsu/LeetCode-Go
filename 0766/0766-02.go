func isToeplitzMatrix(matrix [][]int) bool {
    for r, row := range matrix {
        for c, v := range row {
            if r > 0 && c > 0 && v != matrix[r - 1][c - 1] {
                return false
            }
        }
    }
    
    return true
}
