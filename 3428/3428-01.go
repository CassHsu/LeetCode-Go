func finalPositionOfSnake(n int, commands []string) int {
    x := 0
    d := make([][]int, n)

    for i := range d {
        d[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            d[i][j] = x
            x++
        }
    }

    i, j := 0, 0
    for _, c := range commands {
        switch c {
        case "RIGHT":
            j++
        case "DOWN":
            i++
        case "LEFT":
            j--
        case "UP":
            i--
        }
    }

    return d[i][j]
}
