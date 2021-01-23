import (
	"strconv"
)

func fizzBuzz(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		im3 := i % 3
		im5 := i % 5
		if im3 == 0 && im5 == 0 {
			res = append(res, "FizzBuzz")
		} else if im3 == 0 {
			res = append(res, "Fizz")
		} else if im5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}