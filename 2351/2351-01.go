func repeatedCharacter(s string) byte {
    cs := map[rune]int{}
    r := 0
    
    for i, c := range s {
        if _, ok := cs[c]; ok {
            r = i
            break
        }
        cs[c] = 1
    }
    
    return s[r]
}
