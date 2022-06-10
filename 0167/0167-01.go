func twoSum(numbers []int, target int) []int {
    ans := []int{-1, -1}
    
    l := 0
    r := len(numbers) - 1
    
    for l < r {
        sum := numbers[l] + numbers[r]
        
        if sum == target {
            ans[0] = l + 1
            ans[1] = r + 1
            break
        } else if sum < target {
            l++
        } else {
            r--
        }
    }
    
    return ans
}
