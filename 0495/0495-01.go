func findPoisonedDuration(timeSeries []int, duration int) int {
    count := 0
    
    for i, t := range timeSeries {
        count += duration
        
        if i > 0 {
            diff := t - timeSeries[i - 1]
            if diff < duration {
                count -= (duration - diff)
            }
        }
    }
    
    return count
}
