func minOperations(nums []int) int {
    count := 0
    prev := 0
    
    for _, n := range nums {
        if prev >= n {
            prev++
            count += prev - n
        } else {
            prev = n
        }
    }
    return count
}
