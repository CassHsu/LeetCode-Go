func isPowerOfTwo(n int) bool {
    p := 1
    for i := 0; i < 31; i++ {
        if p == n {
            return true
        }
        p *= 2
    }
    
    return false
}