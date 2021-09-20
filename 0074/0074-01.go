func searchMatrix(matrix [][]int, target int) bool {
    w := len(matrix[0])
    h := len(matrix)
    
    if target < matrix[0][0] { return false }
    if target > matrix[h - 1][w - 1] { return false }
    
    init := make([]int, h)
    for i, m := range matrix {
        init[i] = m[0]
    }
    
    idx := 0
    for i := 0; i < h; i++ {
        if i == h - 1 && init[i] <= target { 
            idx = i
            break
            
        } else {
            if init[i] <= target && target < init[i + 1] {
                idx = i
                break
            }
        }
    }
    
    for _, v := range matrix[idx] {
        if v == target { return true }
    }
    
    return false
}
