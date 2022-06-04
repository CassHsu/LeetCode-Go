func percentageLetter(s string, letter byte) int {
    count := 0
    for _, c := range s {
        if byte(c) == letter {
            count++
        }
    }
    
    ans := float64(count) / float64(len(s)) * 100
    
    return int(ans)
}
