import (
    "strconv"
)

func calPoints(ops []string) int {
    stack := []int{}
    
    for _, op := range ops {
        n := len(stack)
        
        if op == "+" {
            stack = append(stack, stack[n- 1] + stack[n - 2])
            
        } else if op == "C" {
            stack = stack[:n - 1]
            
        } else if op == "D" {
            stack = append(stack, 2 * stack[n - 1])
            
        } else {
            x, _ := strconv.Atoi(op)
            stack = append(stack, x)
        }
    }
    
    ans := 0
    for _, s := range stack {
        ans += s
    }
    return ans
}
