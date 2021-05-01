func squareIsWhite(coordinates string) bool {
    x := (coordinates[0] - 'a' + 1) % 2
    y := (coordinates[1] - '0') % 2
    
    return (x ^ y) == 1
}
