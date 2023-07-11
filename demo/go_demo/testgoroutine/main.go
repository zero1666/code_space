package main

import (
	"fmt"
	"sync"
	"time"
)

var stop = make(chan string)

func main() {

	fmt.Println("vim-go")
}

func testSelectTimeout() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case action := <-stop:
			fmt.Println(action, " gorouting")
		}
	}()

	time.AfterFunc(
		time.Second*2, func() {
			stop <- "stop "
		})
	wg.Wait()

}
