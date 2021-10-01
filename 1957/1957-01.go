func makeFancyString(s string) string {
    ans := ""
    prev := ""
    count := 1
    
    for _, c := range s {
        sc := string(c)
        if prev == sc {
            count++
        } else {
            count = 1
        }
        
        if count < 3 {
            ans += sc
        }
        
        prev = sc
    }
    
    return ans
}
