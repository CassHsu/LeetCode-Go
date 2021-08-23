func transpose(matrix [][]int) [][]int {
    w := len(matrix[0])
    h := len(matrix)
    ans := make([][]int, w)
    
    for c := 0; c < w; c++ {
        ans[c] = make([]int, h)
        for r := 0; r < h; r++ {
            ans[c][r] = matrix[r][c]
        }
    }
    
    return ans
}
