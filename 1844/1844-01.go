func replaceDigits(s string) string {
    ret := ""
    for i := 0; i < len(s); i++ {
        if i % 2 != 0 {
            ret += string(shift(s[i-1], s[i]))
        } else {
            ret += string(s[i])
        }
    }
    return ret
}

func shift(c byte, x byte) string {
    return string(c - '0' + x)
}
