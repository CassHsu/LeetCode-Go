func smallestEqual(nums []int) int {
    for i, n := range nums {
        if n == i % 10 {
            return i
        }
    }
    return -1
}
