func findMiddleIndex(nums []int) int {
    for i, _ := range nums {
        l := sum(nums[:i])
        r := sum(nums[i+1:])
        
        if l == r {
            return i
        }
    }
    
    return -1
}

func sum(nums []int) int {
    sum := 0
    
    for _, s := range nums {
        sum += s
    }
    
    return sum
}
