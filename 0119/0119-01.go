func getRow(rowIndex int) []int {
    ret := [][]int{}
    ret = append(ret, []int{1})
    
    for i := 1; i <= rowIndex; i++ {
        row := make([]int, i+1)
        
        row[0] = 1
        for j := 1; j < i; j++ {
            row[j] = ret[i-1][j-1] + ret[i-1][j]
        }
        row[i] = 1
        ret = append(ret, row)
    }
    return ret[rowIndex]
}
