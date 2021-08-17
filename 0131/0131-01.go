func partition(s string) [][]string {
    pp := pp{s: s, ans: [][]string{}}
    pp.backtrack([]string{}, 0)
    return pp.ans
}

type pp struct {
    s string
    ans [][]string
}

func (p *pp) backtrack(curr []string, start int) {
    if start >= len(p.s) {
        tmp := make([]string, len(curr))
        copy(tmp, curr)
        p.ans = append(p.ans, tmp)
        return
    }
    
    for end := start; end < len(p.s); end++ {
        subStr := p.s[start: end + 1]
        if isPalindrome(subStr) {
            curr = append(curr, subStr)
            p.backtrack(curr, end + 1)
            curr = curr[:len(curr) - 1]
        }
    }
}

func isPalindrome(str string) bool {
    i := 0
    j := len(str) - 1
    
    for i < j {
        if str[i] != str[j] {
            return false
        }
        i++
        j--
    }
    
    return true
}
