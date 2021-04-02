func partitionLabels(S string) []int {
    last := make(map[byte]int)
    for i := 0; i < len(S); i++ {
        last[S[i]] = i;
    }
    
    anchor := 0
    j := 0
    ans := []int{}
    
    for i := 0; i < len(S); i++ {
        if last[S[i]] > j {
            j = last[S[i]]
        }
        if i == j {
            ans = append(ans, i - anchor + 1)
            anchor = i + 1
        }
    }
    return ans
}
