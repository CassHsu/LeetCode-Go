func minimumSum(num int) int {
    ns := strconv.Itoa(num)
    nums := strings.Split(ns, "")
    sort.Strings(nums)
    
    n1, _ := strconv.Atoi(nums[0] + nums[2])
    n2, _ := strconv.Atoi(nums[1] + nums[3])
    
    return n1 + n2
}
