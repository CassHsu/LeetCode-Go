func findPoisonedDuration(timeSeries []int, duration int) int {
    n := len(timeSeries)
    if n == 0 {
        return 0
    }
    
    count := 0
    for i := 0; i < n - 1; i++ {
        diff := timeSeries[i + 1] - timeSeries[i]
        if diff < duration {
            count += diff
        } else {
            count += duration
        }
    }
    
    return count + duration
}
