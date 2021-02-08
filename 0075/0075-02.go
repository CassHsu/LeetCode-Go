func sortColors(nums []int)  {
    count := []int{0, 0, 0}
    
    for i, _ := range nums {
        count[nums[i]] += 1
    }
    
    for i, _ := range nums {
        if count[0] > 0 {
            nums[i] = 0
            count[0]--
            
        } else if (count[1] > 0) {
            nums[i] = 1
            count[1]--
            
        }  else {
            nums[i] = 2
        }
    }
}