func countNegatives(grid [][]int) int {
    count := 0

    for _, row := range grid {
        for _, n := range row {
            if n < 0 {
                count++
            }
        }
    }

    return count
}
