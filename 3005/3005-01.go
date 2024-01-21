func maxFrequencyElements(nums []int) int {
    var m [101]int
    max := 0
    ans := 0

    for _, n := range nums {
        m[n]++

        if m[n] == max {
            ans += max
        } else if m[n] > max {
            max = m[n]
            ans = max
        }
    }

    return ans
}
