func canPlaceFlowers(flowerbed []int, n int) bool {
    i := 0
    size := len(flowerbed)

    for i < size && n > 0 {
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i - 1] == 0) && 
        (i == size - 1 || flowerbed[i + 1] == 0) {
            flowerbed[i] = 1
            n--
        }
        i++
    }

    return n == 0
}
