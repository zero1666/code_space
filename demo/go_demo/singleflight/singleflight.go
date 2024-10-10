package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("exec singleFlightDo...")
	singleFlightDo()
	time.Sleep(time.Second * 5)

	fmt.Println("exec singleFlightDoChan..")
	singleFlightDoChan()
}

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func singleFlightDo() {
	var key string = "http://www.google.com/_do"
	const n = 5
	waited := int32(n)

	doneChan := make(chan struct{})

	var g singleflight.Group

	for i := 0; i < n; i++ {
		go func(j int) {

			val, _, shared := g.Do(key, func() (any, error) {
				// 增加请求量，降低单并发失败带来的影响
				go func() {
					time.Sleep(10 * time.Millisecond)
					fmt.Printf("Deleting key: %v\n", key)
					g.Forget(key)
				}()
				val, err := find(context.Background(), key)
				return val, err
			})

			if atomic.AddInt32(&waited, -1) == 0 {
				close(doneChan)
			}
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, val, shared)
		}(i)
	}

	select {
	case <-doneChan:
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}

func singleFlightDoChan() {
	var key string = "http://www.google.com/do_chan"
	const n = 5

	var g singleflight.Group

	for i := 0; i < n; i++ {
		go func(j int) {

			ch := g.DoChan(key, func() (any, error) {
				val, err := find(context.Background(), key)
				return val, err
			})

			var result singleflight.Result
			// 实现超时控制
			select {
			case result = <-ch:
				fmt.Printf("index: %d, val: %q, shared: %v\n", j, result.Val.(Result), result.Shared)
			case <-time.After(time.Millisecond * 5):
				fmt.Printf("index:%d, timeout \n", j)
			}

		}(i)
	}

	time.Sleep(time.Second * 5)
}
