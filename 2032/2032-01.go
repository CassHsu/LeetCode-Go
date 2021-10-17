type NOutOfM struct {
    m map[int]int
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    tt := NOutOfM { m: map[int]int{} }
    
    tt.Set(nums1)
    tt.Set(nums2)
    tt.Set(nums3)
    
    return tt.Count(2)
}

func (t *NOutOfM) Count(n int) []int {
    ret := []int{}
    for k, v := range t.m {
        if v >= n {
            ret = append(ret, k)
        }
    }
    return ret
}

func (t *NOutOfM) IntoMap(nums []int) {
    for _, n := range nums {
        _, ok := t.m[n]
        if !ok {
            t.m[n] = 1
        } else {
            t.m[n]++
        }
    }
}

func (t *NOutOfM) Set(nums []int) {
    ret := []int{}
    m := map[int]int{}
    
    for _, n := range nums {
        _, ok := m[n]
        if !ok {
            m[n] = 1
            ret = append(ret, n)
        }
    }
    t.IntoMap(ret)
}
