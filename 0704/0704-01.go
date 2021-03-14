func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            return mid
        }
        if target < nums[mid] {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return -1
}
