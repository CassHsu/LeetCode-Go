func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }
    
    m := [26]int{}
    for _, s := range sentence {
        m[s - 'a'] += 1
    }
    
    for _, v := range m {
        if v == 0 {
            return false
        }
    }
    
    return true
}
