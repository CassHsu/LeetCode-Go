func daysBetweenDates(date1 string, date2 string) int {
    const date_format = "2006-01-02"
    d1, _ := time.Parse(date_format, date1)
    d2, _ := time.Parse(date_format, date2)
    
    days := int(d2.Sub(d1).Hours()) / 24
    if days > 0 {
        return days
    }
    return days * -1
}
