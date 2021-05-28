func findMaxConsecutiveOnes(nums []int) int {
    max1 := 0
    count := 0
    
    for _, n := range nums {
        if n == 1 {
            count++
        } else {
            count = 0
        }
        
        if count > max1 {
            max1 = count
        }
    }
    return max1
}
