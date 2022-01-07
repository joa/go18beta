# Go 1.18beta test repo
Experimenting with Go 1.18beta generics.

### Future Example
Create and return an asynchronous computation.
```go
func asyncFib(n int) future.Future[int] {
    return future.Go[int](func() (fib1 int, err error) {
        if n == 0 {
            return
        }
        
        fib0 := 0
        fib1 = 1
        
        for i := 2; i <= n; i++ {
            fib2 := fib0 + fib1
            fib0 = fib1
            fib1 = fib2
        }
        
        return
    })
}

// Start the async computation and do something with
// the value once ready.
// All callbacks are executed in their own go routine.
future.Map(asyncFib(n), func(res int) string {
    // Map the integer to a string.
    return strconv.Itoa(res)
}).Then(func(res string) {
    // Listen for completion of the mapped future.
    fmt.Println(res)
}).Catch(func(err error) {
    // Handle the error case.
    panic(err)
})
```

### Promise/Future Example
More advanced use case that deals with the read- and write-only side
of the asynchronous computation. This is useful when completion of the promise is done by multiple
competing producers for instance.
```go
func asyncFib(n int) future.Future[int] {
    // Create a promise
    // This is the write-only side of an async computation
    res := future.Create[int]()

    go func() {
        if n == 0 {
            res.Resolve(0) 
            return
        }

        fib0 := 0
        fib1 := 1

        for i := 2; i <= n; i++ {
            fib2 := fib0 + fib1
            fib0 = fib1
            fib1 = fib2
        }

        // Resolve the promise
        res.Success(fib1)
    }()

    // Return the Promise's future.
    // This is the read-only side of the async computation.
    return res.Future() 
}
```