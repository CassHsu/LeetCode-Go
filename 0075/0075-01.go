func sortColors(nums []int)  {
    size := len(nums)
    stop := 1
    isConti := true
    for isConti {
        isConti = false
        count := size - stop
        for i := 0; i < count; i++ {
            if nums[i] > nums[i+1] {
                nums[i], nums[i+1] = nums[i+1], nums[i]
                isConti = true
            }
        }
        stop++
    }
}