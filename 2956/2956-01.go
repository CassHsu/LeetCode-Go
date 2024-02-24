func findIntersectionValues(nums1 []int, nums2 []int) []int {
    m := make(map[int]bool)
    for _, n := range nums1 {
        if _, ok := m[n]; !ok {
            m[n] = true
        }
    }

    c1, c2 := 0, 0
    for _, n := range nums2 {
        if m[n] == true {
            c1 += count(nums1, n)
            c2 += count(nums2, n)
            m[n] = false
        }
    }

    return []int{c1, c2}
}

func count(arr []int, target int) int {
    cnt := 0
    for _, n := range arr {
        if n == target {
            cnt++
        }
    }
    return cnt
}
