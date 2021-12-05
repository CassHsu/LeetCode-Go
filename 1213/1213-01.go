import (
    "sort"
)

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    m := map[int]int{}
    ans := []int{}
    
    countIt(&arr1, &m)
    countIt(&arr2, &m)
    countIt(&arr3, &m)
    
    for k, v := range m {
        if v == 3 {
            ans = append(ans, k)
        }
    }
    
    sort.Ints(ans)
    return ans
}

func countIt(arr *[]int, m *map[int]int) {
    for _, n := range *arr {
        c, ok := (*m)[n]
        if ok {
            (*m)[n] = c + 1
        } else {
            (*m)[n] = 1
        }
    }
}
