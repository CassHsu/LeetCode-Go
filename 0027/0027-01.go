func removeElement(nums []int, val int) int {
    i := 0
    for j, _ := range nums {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}
