func simplifyPath(path string) string {
    stack := []string{}
    s := strings.Split(path, "/")
    
    for _, c := range s {
        if c == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
            
        } else if c != "" && c != "." {
            stack = append(stack, c)
        }
    }
    
    return "/" + strings.Join(stack, "/")
}
