func finalValueAfterOperations(operations []string) int {
    ans := 0
    
    for _, op := range operations {
        if op == "--X" || op == "X--" { ans-- }
        if op == "++X" || op == "X++" { ans++ }
    }
    
    return ans
}
