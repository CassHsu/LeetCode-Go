import (
    "strings"
)

func numUniqueEmails(emails []string) int {
    m := map[string]int{}
    
    for _, email := range emails {
        names := strings.Split(email, "@")
        
        local := strings.Split(names[0], "+")[0]
        locals := strings.Split(local, ".")
        local = strings.Join(locals, "")
        
        k := local + "@" + names[1]
        _, ok := m[k]
        if !ok {
            m[k] = 1
        }
    }
    
    return len(m)
}
