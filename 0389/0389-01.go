func findTheDifference(s string, t string) byte {
    st := s + t
    r := byte(st[0])
    for i := 1; i < len(st); i++ {
        r ^= byte(st[i])
    }
    return r
}
