import (
    "strings"
    "strconv"
)

func validIPAddress(IP string) string {
    if vaildIPv4(IP) == true {
        return "IPv4"
    }
    if vaildIPv6(IP) == true {
        return "IPv6"
    }
    return "Neither"
}

func vaildIPv4(ip string) bool {
    seg := strings.Split(ip, ".")
    if len(seg) != 4 {
        return false
    }
    
    for _, s := range seg {
        if vaildIPv4Seg(s) == false {
            return false
        }
    }
    return true
}

func vaildIPv4Seg(seg string) bool {
    size := len(seg)
    if size == 0 || size > 3 {
        return false
    }
    if size >= 2 && string(seg[0]) == "0" {
        return false
    }
    n, err := strconv.Atoi(seg)
    if err != nil {
        return false
    }
    if n < 0 || n > 255 {
        return false
    }
    return true
}

func vaildIPv6(ip string) bool {
    seg := strings.Split(ip, ":")
    if len(seg) == 0 || len(seg) > 8 {
        return false
    }
    
    for _, s := range seg {
        if vaildIPv6Seg(s) == false {
            return false
        }
    }
    return true
}

func vaildIPv6Seg(seg string) bool {
    size := len(seg)
    if size == 0 || size > 4 {
        return false
    }
    for _, s := range seg {
        if isHex(s) == false {
            return false
        }
    }
    return true
}

func isHex(s rune) bool {
    for _, x := range "0123456789abcdefABCDEF" {
        if s == x {
            return true
        }
    }
    return false
}
