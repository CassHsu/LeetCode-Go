func minMovesToSeat(seats []int, students []int) int {
    count := 0
    sort.Ints(seats)
    sort.Ints(students)
    
    for i := 0; i < len(seats); i++ {
        diff := seats[i] - students[i]
        if diff < 0 {
            diff *= -1
        }
        count += diff
    }
    
    return count
}
