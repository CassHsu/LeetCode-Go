func getMinDistance(nums []int, target int, start int) int {
    ans := 100000
    for i, n := range(nums) {
        if n == target {
            diff := i - start
            if diff < 0 {
                diff = -diff
            }
            if diff < ans {
                ans = diff
            }
        }
    }
    return ans
}
