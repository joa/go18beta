# Go 1.18beta test repo
Experimenting with Go 1.18beta generics.

### Promise/Future Example
```go

func asyncFib(n int) future.Future[int] {
	res := future.Create[int]() // create a promise, write-only side of async computation

	go func() {
		if n == 0 {
			res.Success(0) // resolve the promise
			return
		}

		fib0 := 0
		fib1 := 1

		for i := 2; i <= n; i++ {
			fib2 := fib0 + fib1
			fib0 = fib1
			fib1 = fib2
		}

		res.Success(fib1) // resolve the promise
	}()

	return res.Future() // return the read-only side of the async computation
}

future.Map(asyncFib(n), func(res int) string {
    return strconv.Itoa(res)
}).Then(func(res string) {
    fmt.Println(res)
}).Catch(func(err error) {
    panic(err)
})
```