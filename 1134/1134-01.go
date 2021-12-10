import (
    "fmt"
    "strconv"
    "math"
)

func isArmstrong(n int) bool {
    total := 0
    sn := strconv.Itoa(n)
    size := len(sn)
    
    for _, c := range sn {
        num, _ := strconv.Atoi(string(c))
        total += int(math.Pow(float64(num), float64(size)))
    }
    
    return total == n
}
