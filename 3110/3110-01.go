func scoreOfString(s string) int {
    score := 0

    for i := 1; i < len(s); i++ {
        score = score + int(math.Abs(float64(int(s[i-1]) - int(s[i]))))
    }

    return score
}
