package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//testCancelContext()
	testSubCancel()
	fmt.Println("vim-go")
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

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("start cancel")
		cancelFunc()
	})
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

	subCtx2, _ := context.WithCancel(subCtx1)
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
