func getSumAbsoluteDifferences(nums []int) []int {
    size := len(nums)
    total := 0
    for _, n := range nums {
        total += n
    }
    result := make([]int, size)
    curr := 0
    
    for i, n := range nums {
        curr += n
        result[i] = (i+1)*n - curr + (total - curr) - (size-i-1)*n 
    }
    
    return result
}
