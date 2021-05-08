func generate(numRows int) [][]int {
    ret := make([][]int, 0)
    
    for i := 0; i < numRows; i++ {
        row := make([]int, i+1)
        row[0] = 1
        
        for j := 1; j < i; j++ {
            row[j] = ret[i-1][j-1] + ret[i-1][j]
        }
        
        row[i] = 1
        ret = append(ret, row)
    }
    return ret
}
