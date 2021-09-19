func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c { return mat }
    
    q := []int{}
    for _, row := range mat {
        for _, v := range row {
            q = append(q, v)
        }
    }
    
    ans := [][]int{}
    k := 0
    for i := 0; i < r; i++ {
        row := []int{}
        for j := 0; j < c; j++ {
            row = append(row, q[k])
            k++
        }
        
        ans = append(ans, row)
    }
    
    return ans
}
