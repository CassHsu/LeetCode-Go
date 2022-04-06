func maxArea(height []int) int {
    ans := 0
    l, r := 0, len(height) - 1
    
    for l < r {
        curr_h := height[l]
        if height[r] < curr_h {
            curr_h = height[r]
        }
        
        area := (r - l) * curr_h
        if area > ans {
            ans = area
        }
        
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    
    return ans
}
