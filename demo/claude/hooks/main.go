package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	testTimeoutControl()
	fmt.Println("vim-go")
	_ = strings.ToUpper("test") // Prevent strings import from being removed
}

func testCancelContext() {
	var wg sync.WaitGroup

	wg.Add(1)
	ctx := context.Background()
	subCtx, cancelFunc := context.WithCancel(ctx)
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("subctx Done")
			return
		}
	}(subCtx)

	wg.Wait()
}

func testSubCancel() {
	var wg sync.WaitGroup
	wg.Add(2)

	ctx := context.Background()
	subCtx1, cancelFunc1 := context.WithCancel(ctx)

	go func() {
		defer wg.Done()
		subFunc(subCtx1, "myfunc")
	}()

	subCtx2, cancelFunc2 := context.WithCancel(subCtx1)
	defer cancelFunc2()
	go func() {
		defer wg.Done()
		subFunc(subCtx2, "mysubfunc")
	}()

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("start cancel1")
		cancelFunc1()
	})
	wg.Wait()
}

func subFunc(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Println(name, " Done")
		return
	}
}

func testPassValue() {
	traceID := "12345"
	ctx := context.WithValue(context.Background(), "traceID", traceID)
	go processRequest(ctx)
	time.Sleep(time.Second * 1)

}
func processRequest(ctx context.Context) {
	traceID, ok := ctx.Value("traceID").(string)
	if ok {
		fmt.Printf("Processing request with traceID: %s\n", traceID)
	} else {
		fmt.Println("Processing request without traceID")
	}
}

func testConcurrencyControl() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range genFibonacci(ctx) {
		fmt.Println(n)
		if n >= 100 {
			cancel()
			break
		}
	}

}

func genFibonacci(ctx context.Context) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		a, b := 0, 1
		for {
			select {
			case <-ctx.Done():
				return
			case c <- a:
				a, b = b, a+b
				//time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	return c
}

func testTimeoutControl() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	retries := 0

	const maxTryNum = 5
	retryInterval := 500 * time.Millisecond
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			retries++
			if retries > maxTryNum {
				fmt.Println("Maximum retries reached")
				return
			}
			fmt.Printf("retrying...%d/%d \n", retries, maxTryNum)
			time.Sleep(retryInterval)

		}
	}

}
