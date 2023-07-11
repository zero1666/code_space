// test.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	test4()
}

func test3() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		buf := make([]byte, 16*1024*1024)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
	}()
	sum := 0
	for {
		n := rand.Intn(1e6)
		sum += n
		if sum%42 == 0 {
			panic(":()")
		}
	}
}
func test2() {
	sum := 0
	for {
		n := rand.Intn(1e6)
		sum += n
		if sum%42 == 0 {
			panic(":()")
		}
	}
}

func test4() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		buf := make([]byte, 16*1024*1024)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
	}()

	for i := 1; i <= 10; i++ {
		go func(gid int) {
			defer func() {
				err := recover()
				if err == nil {
					return
				}
				buf := make([]byte, 16*1024*1024)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
			}()
			n := 0
			for {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), gid, n)
				time.Sleep(time.Second)
			}
		}(i)
	}

	go func() {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			buf := make([]byte, 16*1024*1024)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
		}()
		arr := 0
		p := uintptr(unsafe.Pointer(&arr))
		myfun1(p)
	}()

	for true {
		time.Sleep(time.Second)
	}

}
func test1() {
	for i := 1; i <= 10; i++ {
		go func(gid int) {
			n := 0
			for {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), gid, n)
				time.Sleep(time.Second)
			}
		}(i)
	}

	go func() {
		arr := 0
		p := uintptr(unsafe.Pointer(&arr))
		myfun1(p)
	}()

	for true {
		time.Sleep(time.Second)
	}

}

func myfun1(p uintptr) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		buf := make([]byte, 16*1024*1024)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
	}()
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
	go myfun2()
	fmt.Println(*arr)
}

func myfun2() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		buf := make([]byte, 16*1024*1024)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
	}()
	fmt.Println("myfun2")
	myfun3()
}

func myfun3() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		buf := make([]byte, 16*1024*1024)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("[ES NULL POINT ERROR]%v\n%s\n", err, buf)
	}()
	var p uintptr = 0
	arr := (*int)(unsafe.Pointer(p))
	*arr = 1
	fmt.Println(*arr)
}
