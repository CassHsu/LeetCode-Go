func canPermutePalindrome(s string) bool {
    count := 0
    m := map[rune]int{}
    
    for _, r := range s {
        m[r]++
    }
    
    for _, v := range m {
        count += v % 2
    }
    
    return count <= 1
}
