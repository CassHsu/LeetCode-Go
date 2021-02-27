func checkInclusion(s1 string, s2 string) bool {
    size1 := len(s1)
    size2 := len(s2)
    
    if size1 > size2 {
        return false
    }
    
    arr1 := [26]int{}
    arr2 := [26]int{}
    for i := 0; i < size1; i++ {
        arr1[s1[i] - 'a']++
        arr2[s2[i] - 'a']++
    }
    
    w := size2 - size1
    for i := 0; i < w; i++ {
        if matches(arr1, arr2) {
            return true
        }
        
        arr2[s2[i] - 'a']--
        arr2[s2[i + size1] - 'a']++
    }
    return matches(arr1, arr2)
}

func matches(arr1 [26]int, arr2 [26]int) bool {
    for i := 0; i < 26; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
