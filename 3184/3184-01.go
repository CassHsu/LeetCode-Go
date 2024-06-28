func countCompleteDayPairs(hours []int) int {
    count := 0
    size := len(hours)

    for i := 0; i < size; i++ {
        for j := i + 1; j < size; j++ {
            if (hours[i] + hours[j]) % 24 == 0 {
                count++
            }
        }
    }

    return count
}
