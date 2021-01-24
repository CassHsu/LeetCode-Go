func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := map[int]int{}

	for _, n := range nums1 {
		_, ok := m[n]
		if ok {
			m[n]++
		} else {
			m[n] = 1
		}
	}

	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}

	return res
}