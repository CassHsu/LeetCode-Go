func countPairs(nums []int, target int) int {
    count := 0
    size := len(nums)

    for i := 0; i < size; i++ {
        for j := i + 1; j < size; j++ {
            if (nums[i] + nums[j]) < target {
                count += 1
            }
        } 
    }

    return count
}
