func isUgly(n int) bool {
    if n == 0 {
        return false
    }
    arr := []int {2, 3, 5}
    
    for _, a := range arr {
        for n % a == 0 {
            n /= a
        }
    }
    
    return n == 1
}
