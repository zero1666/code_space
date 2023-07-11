package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 6}

	arr2 := arr[0:100]
	fmt.Printf("a:%v, arr2:%v", arr, arr2)
	fmt.Println("vim-go")
}
