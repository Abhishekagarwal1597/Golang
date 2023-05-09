var cache = make(map[int]int)

func fib(n int) int {
    if n <= 1 {
        return n
    }

    if num, ok := cache[n]; ok {
        return num
    }
    cache[n] = fib(n-1) + fib(n-2)

    return cache[n]
}
