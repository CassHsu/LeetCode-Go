func countPoints(rings string) int {
    count := 0
    m := map[byte]string {}
    
    for r, c := 0, 1; c < len(rings); r, c = r + 2, c + 2 {
        k := rings[c]
        v, ok := m[k]
        if !ok { 
            m[k] = ""
        }
        m[k] = v + string(rings[r])
    }
    
    for _, v := range m {
        if len(v) < 3 { continue }
        
        r := false
        g := false
        b := false
        
        for _, c := range v {
            if string(c) == "R" { r = true }
            if string(c) == "G" { g = true }
            if string(c) == "B" { b = true }
            
            if r && g && b {
                count++
                break
            }
        }
    }
    
    return count
}
