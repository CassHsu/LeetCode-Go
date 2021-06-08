func rotate(nums []int, k int)  {
    size := len(nums)
    k = k % size
    
    reverse(nums[:size - k])
    reverse(nums[size- k:])
    reverse(nums)
    
}

func reverse(nums []int) {
    l := 0
    r := len(nums) - 1
    
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        
        l++
        r--
    }
}
