func replaceElements(arr []int) []int {
    max := arr[len(arr) - 1]
    arr[len(arr) - 1] = -1
    
    for i := len(arr) - 2; i >= 0; i-- {
        tmp := max
        if arr[i] > max {
            max = arr[i]
        }
        arr[i] = tmp
    }
    return arr
}
