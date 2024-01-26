func maxFrequencyElements(nums []int) int {
    m := map[int]int{}
    max := 0
    count := 0

    for _, n := range nums {
        f := m[n] + 1
        if f == max {
            count += 1
        } else if f > max {
            max = f
            count = 1
        }
        m[n] = f
    }

    return max * count
}
