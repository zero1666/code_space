package main

import "fmt"

type A struct {
	v1 int
	v2 string
}

type B struct {
	*A
	v3 int
}

func main() {
	var a = A{v1: 1, v2: "v2"}
	var b B
	b.v3 = 3
	b.A = &a
	fmt.Printf("a:%v, \n b:%v\n", a, b)
}
