func rotateString(s string, goal string) bool {
    if len(s) == 0 && len(goal) == 0 {
        return true
    }
    
    if len(s) != len(goal) {
        return false
    }
    
    if s == goal {
        return true
    }
    
    g := goal
    for i := 0; i < len(s); i++ {
        g = g[1:] + g[:1]
        if s == g {
            return true
        }
    }
    return false
}
