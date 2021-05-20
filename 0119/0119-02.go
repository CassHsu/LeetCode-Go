func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    if rowIndex == 1 {
        return []int{1, 1}
    }
    
    last := getRow(rowIndex - 1)
    row := make([]int, len(last)+1)
    row[0] = 1
    row[rowIndex] = 1
    
    for j := 1; j < len(last); j++ {
        row[j] = last[j-1] + last[j]   
    }
    return row
}
