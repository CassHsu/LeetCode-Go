func largestLocal(grid [][]int) [][]int {
    res := [][]int{}
    
    for r := 0; r <= len(grid) - 3; r++ {
        mx := []int{}
        for c := 0; c <= len(grid[0]) - 3; c++ {
            v1 := max3(grid[r][c], grid[r][c + 1], grid[r][c + 2])
            v2 := max3(grid[r + 1][c], grid[r + 1][c + 1], grid[r + 1][c + 2])
            v3 := max3(grid[r + 2][c], grid[r + 2][c + 1], grid[r + 2][c + 2])
            
            mx = append(mx, max3(v1, v2, v3))
        }
        
        res = append(res, mx)
    }
    
    return res
}

func max3(a int, b int, c int) int {
    if a >= b && a >= c { return a }
    if b >= a && b >= c { return b }
    if c >= a && c >= b { return c }
    return -1
}
