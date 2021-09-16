type Fib struct {
    memo map[int]int
}

func fib(n int) int {
    f := Fib { memo: map[int]int  { 0: 0, 1: 1 } }
    return f.fib(n)
}

func (f *Fib) fib(n int) int {
    v, ok := f.memo[n]
    if ok { return v }
    
    v = f.fib(n - 1) + f.fib(n - 2)
    f.memo[n] = v
    return v
}
