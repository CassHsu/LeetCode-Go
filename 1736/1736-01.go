import (
    "strconv"
)

func maximumTime(time string) string {
    res := []byte(time)
    
    if res[0] == '?' {
        t1, err := strconv.Atoi(string(res[1]))
        if err == nil && t1 > 3 {
            res[0] = '1'
        } else {
            res[0] = '2'
        }
    }
    if res[1] == '?' {
        if res[0] == '2' {
            res[1] = '3'
        } else {
            res[1] = '9'
        }
    }
    if res[3] == '?' { res[3] = '5' }
    if res[4] == '?' { res[4] = '9' }

    return string(res)
}