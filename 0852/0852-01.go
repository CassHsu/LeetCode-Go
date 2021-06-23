func peakIndexInMountainArray(arr []int) int {
    pi := 0
    pv := arr[0]
    
    for i := 1; i < len(arr); i++ {
        if arr[i] > pv {
            pi = i
            pv = arr[i]
        }
    }
    
    return pi
}
