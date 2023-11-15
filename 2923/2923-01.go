func findChampion(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    for i := 0; i < n; i++ {
      sum := 0
      for j := 0; j < m; j++ {
        sum += grid[i][j]
      }

      if sum == n - 1 {
        return i
      }
    }
    
    return n
}
