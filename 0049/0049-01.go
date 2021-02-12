import (
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    ret := [][]string{}
    if len(strs) == 0 {
        return ret;
    }
    
    m := map[string][]string{}
    for _, s := range strs {
        sp := strings.Split(s, "")
        sort.Strings(sp)
        k := strings.Join(sp, "")
        
        m[k] = append(m[k], s)
    }
    
    for _, v := range m {
        ret = append(ret, v)
    }
    return ret
}