func guessNumber(n int) int {
    mid := n / 2
    res := guess(mid)
    
    for true {
        if res == 0 {
            return mid
        } else if res > 0 {
            mid = (mid + 1 + n) / 2
            res = guess(mid)
        } else {
            n = mid - 1
            mid = (1 + n) / 2
            res = guess(mid)
        }
    }
    
    return -1
}
