import (
    "sort"
    "fmt"
)

func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
    
    sort.Sort(sort.Reverse(sort.IntSlice(stones)))
    
    m1 := stones[0]
    m2 := stones[1]
    stones = stones[2:]
    stones = append(stones, m1 - m2)
    
    return lastStoneWeight(stones)
}
