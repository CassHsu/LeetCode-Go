var L = byte('1')
var W = byte('0')

func numIslands(grid [][]byte) int {
    ROW := len(grid)
    if ROW <= 0 {
        return 0
    }
    
    COL := len(grid[0])
    islands := 0
    
    for i, r := range grid {
        for j, g := range r {
            if g == L {
                islands++
                dfs(grid, ROW, COL, i, j)
            }
        }
    }
    
    return islands
}

func dfs(grid [][]byte, row int, col int, r int, c int) {
    if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] != L {
        return
    }
    
    grid[r][c] = W
    dfs(grid, row, col, r+1, c)
    dfs(grid, row, col, r, c+1)
    dfs(grid, row, col, r-1, c)
    dfs(grid, row, col, r, c-1)
}