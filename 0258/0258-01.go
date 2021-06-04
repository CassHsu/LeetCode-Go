func addDigits(num int) int {
    if num == 0 {
        return 0
    }
    
    for num >= 10 {
        num = num / 10 + num % 10
    }
    return num
}
