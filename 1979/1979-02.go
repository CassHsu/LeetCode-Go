func findGCD(nums []int) int {
    small := nums[0]
    large := nums[0]
    
    for _, n := range nums {
        if n < small {
            small = n
        } else if n > large {
            large = n
        }
    }
    
    return gcd(large, small)
}

func gcd(a int, b int) int {
    if b == 0 {
        return a
    }
    
    return gcd(b, a % b)
}
