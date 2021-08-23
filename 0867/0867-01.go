func transpose(matrix [][]int) [][]int {
    ans := [][]int{}
    w := len(matrix[0])
    h := len(matrix)
    
    for c := 0; c < w; c++ {
        tmp := []int{}
        for r := 0; r < h; r++ {
            tmp = append(tmp, matrix[r][c])
        }
        
        ans = append(ans, tmp)
    }
    
    return ans
}
