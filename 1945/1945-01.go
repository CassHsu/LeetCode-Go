import (
    "strconv"
)

func getLucky(s string, k int) int {
    ts := convert(s)
    
    for k > 0 {
        ts = transform(ts)
        k--
    }
    
    ans, _ := strconv.Atoi(ts)
    return ans
}

func convert(s string) string {
    ns := ""
    for _, c := range s {
        ns += strconv.Itoa(int(c) - 96)
    }
    return ns
}

func transform(s string) string {
    sum := 0
    for _, c := range s {
        nc, _ := strconv.Atoi(string(c))
        sum += nc
    }
    return strconv.Itoa(sum)
}
