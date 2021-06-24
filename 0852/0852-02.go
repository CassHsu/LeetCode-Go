func peakIndexInMountainArray(arr []int) int {
    l := 0
    h := len(arr) - 1
    
    for l < h {
        m := (l + h) / 2
        if arr[m] <  arr[m + 1] {
            l = m + 1
        } else {
            h = m
        }
    }
    
    return l
}
