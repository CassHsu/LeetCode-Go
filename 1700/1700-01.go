func countStudents(students []int, sandwiches []int) int {
    if count(students, 1) == count(sandwiches, 1) {
        return 0
    }
    
    for len(sandwiches) > 0 {
        if count(students, sandwiches[0]) == 0 {
            break
        }
        
        student := students[0]
        students = students[1:]
        if sandwiches[0] == student {
            sandwiches = sandwiches[1:]
        } else {
            students = append(students, student)
        }
    }
    return len(students)
}

func count(arr []int, target int) int {
    count := 0
    for _, a := range arr {
        if a == target {
            count++
        }   
    }
    return count
}