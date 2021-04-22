func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }
    
    m := map[rune]int{}
    for _, c := range "abcdefghijklmnopqrstuvwxyz" {
        m[c] = 0
    }
    
    for _, s := range sentence {
        m[s]++
    }
    
    for _, v := range m {
        if v == 0 {
            return false
        }
    }
    
    return true
}
