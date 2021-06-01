func isIsomorphic(s string, t string) bool {
    ms := map[byte]int{}
    mt := map[byte]int{}
    
    for i := 0; i < len(s); i++ {        
        sv, s_ok := ms[s[i]]
        tv, t_ok := mt[t[i]]
        if s_ok != t_ok || sv != tv {
            return false
        }
        
        ms[s[i]] = i
        mt[t[i]] = i
    }    
    return true
}
