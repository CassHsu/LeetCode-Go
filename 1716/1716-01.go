func totalMoney(n int) int {
    total := 0
    w := int(n / 7)
    r := int(n % 7)
    
    monday := 1
    for i := 0; i < w; i++ {
        t := monday
        for j := 0; j < 7; j++ {
            total += t
            t++
        } 
        monday++
    }
    
    t := monday
    for i := 0; i < r; i++ {
        total += t
        t++
    }
    
    return total
}