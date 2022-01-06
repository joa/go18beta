# ｇｏ ｇｅｎｅｒｉｃｓ ｓｕｐｅｒ ｄｕｐｅｒ ｔｅｓｔ ｒｅｐｏ

![yey generics](https://media.giphy.com/media/3oEduTny9qJEtpGElG/giphy.gif)

much generics such wow 

### 乇ﾒﾑﾶｱﾚ乇
```go

func asyncFib(n int) future.Future[int] {
	res := future.Create[int]()

	go func() {
		if n == 0 {
			res.Success(0)
			return
		}

		fib0 := 0
		fib1 := 1

		for i := 2; i <= n; i++ {
			fib2 := fib0 + fib1
			fib0 = fib1
			fib1 = fib2
		}

		res.Success(fib1)
	}()

	return res.Future()
}

future.Map(asyncFib(n), func(res int) string {
    return strconv.Itoa(res)
}).Then(func(res string) {
    fmt.Println(res)
}).Catch(func(err error) {
    panic(err)
})
```