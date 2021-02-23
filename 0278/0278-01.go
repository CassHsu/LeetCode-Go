func firstBadVersion(n int) int {
    start := 1
    end := n
    
    for true {
        mid := start + (end - start) / 2
        if isBadVersion(mid) && !isBadVersion(mid - 1) {
            return mid
        }
        
        if isBadVersion(mid) == false {
            start = mid + 1
        } else {
            end = mid
        }
    }
    
    return 0
}
