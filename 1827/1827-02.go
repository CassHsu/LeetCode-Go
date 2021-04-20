func minOperations(nums []int) int {
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] >= nums[i] {
            diff := nums[i-1] - nums[i]
            diff++
            count += diff
            nums[i] += diff
        }
    }
    return count
}
