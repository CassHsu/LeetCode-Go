func findPermutationDifference(s string, t string) int {
    ans := 0

    for i, c := range s {
        d := i - strings.Index(t, string(c))
        
        if d < 0 {
            ans -= d
        } else {
            ans += d
        }
    }

    return ans
}
