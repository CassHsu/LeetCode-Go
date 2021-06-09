func replaceElements(arr []int) []int {
    ret := []int{}
    
    for i := 0; i < len(arr) - 1; i++ {
        max := arr[i+1]
        for j := i + 2; j < len(arr); j++ {
            if arr[j] > max {
                max = arr[j]
            }
        }
        ret = append(ret, max)
    }
    
    ret = append(ret, -1)
    return ret
}
