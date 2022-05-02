func backspaceCompare(s string, t string) bool {
    return buildit(s) == buildit(t)
}

func buildit(st string) string {
    r := make([]rune, 0)
    
    for _, c := range st {
        if c == '#' {
            if len(r) > 0 {
                r = r[:len(r) - 1]
            }
            
        } else {
            r = append(r, c)
        }
    }
    
    return string(r)
}
