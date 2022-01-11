func checkString(s string) bool {
    a := 0
    b := len(s)
    
    for i, c := range s {
        if string(c) == "a" {
            a = i
        } else if b == len(s) {
            b = i
        }
    }
    
    return a <= b
}
