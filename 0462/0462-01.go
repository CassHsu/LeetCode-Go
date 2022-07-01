import (
    "sort"
)

func minMoves2(nums []int) int {
    sort.Ints(nums)
    p := nums[len(nums) / 2]
    
    moves := 0
    for _, n := range nums {
        diff := n - p
        if diff > 0 {
            moves += diff
        } else if diff < 0 {
            moves -= diff
        }
    }
    
    return moves
}
