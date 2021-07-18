func checkPerfectNumber(num int) bool {
    if num <= 0 {
        return false
    }
    
    sum := 0
    half := num >> 1
    for i := 1; i <= half; i++ {
        if num % i == 0 {
            sum += i
        }
    }
    
    return sum == num
}
