func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    
    m := make(map[rune]rune)
    
    m['('] = ')'
    m['['] = ']'
    m['{'] = '}'
    
    stack := make([]rune, 0)
    
    for _, c := range s {
        size := len(stack)
        if size > 0 && m[stack[size - 1]] == c {
            stack = stack[:size - 1]
        } else {
            stack = append(stack, c)
        }
    }
    
    return len(stack) == 0
}
