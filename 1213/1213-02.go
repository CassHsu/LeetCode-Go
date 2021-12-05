func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    ans := []int{}
    p1 := 0
    p2 := 0
    p3 := 0
    
    for p1 < len(arr1) && p2 < len(arr2) && p3 < len(arr3) {
        if arr1[p1] == arr2[p2] && arr2[p2] == arr3[p3] {
            ans = append(ans, arr1[p1])
            p1++
            p2++
            p3++
            
        } else {
            if arr1[p1] < arr2[p2] {
                p1++
            } else if arr2[p2] < arr3[p3] {
                p2++
            } else {
                p3++
            }
        }
    }
    
    return ans
}
