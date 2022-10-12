func countAsterisks(s string) int {
    count := 0
    bars := 0
    
    for _, c := range s {
        switch (c) {
            case '*':
            if bars % 2 == 0 {count++}
            break
            
            case '|':
            bars++
            break
            
            default:
            break
        }
    }
    
    return count
}
