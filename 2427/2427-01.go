func commonFactors(a int, b int) int {
    count := 1
    min := b
    if b >= a {
        min = a
    }
    
    for n := 2; n <= min; n++ {
        if a % n == 0 && b % n == 0 {
            count++
        }
    }
    
    return count
}
