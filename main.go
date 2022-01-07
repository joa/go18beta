package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joa/go18beta/future"
)

func asyncFibV1(n int) future.Future[int] {
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

func asyncFibV2(n int) future.Future[int] {
	res := future.Create[int]() // create write only side (promise)

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

		res.Resolve(fib1)
	}()

	return res.Future() // return readonly side (future)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [n]\n", os.Args[0])
		os.Exit(1)
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("usage: %s [n]\n", os.Args[0])
		os.Exit(1)
		return
	}

	exit := make(chan bool)

	future.Map(asyncFibV1(n), func(fib int) string {
		// conversion to string
		return strconv.Itoa(fib)
	}).Then(func(fib string) {
		// handle result of computation
		fmt.Println(fib)
		exit <- true
	}).Catch(func(err error) {
		// deal with error
		fmt.Println("yolo")
		exit <- true
	})

	select {
	case <-sig():
	case <-exit:
	}
}

func sig() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT)
	return sig
}
