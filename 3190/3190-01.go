func minimumOperations(nums []int) int {
    count := 0

    for _, n := range nums {
        if n % 3 != 0 {
            count += 1
        }
    }
    
    return count
