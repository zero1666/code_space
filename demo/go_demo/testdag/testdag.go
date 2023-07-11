package main

import (
	"fmt"
	"learngo/dag"
)

func main() {
	d := dag.New()
	d.Spawns(f1, f2, f3).Join().Pipeline(f4, f5).Then().Spawns(f6, f7, f8)
	d.Run()

	fmt.Println("vim-go")
}

func f1() {
	fmt.Println("run f1")
}

func f2() {
	fmt.Println("run f2")
}

func f3() {
	fmt.Println("run f3")
}

func f4() {
	fmt.Println("run f4")
}

func f5() {
	fmt.Println("run f5")
}

func f6() {
	fmt.Println("run f6")
}

func f7() {
	fmt.Println("run f7")
}

func f8() {
	fmt.Println("run f8")
}

func f9() {
	fmt.Println("run f9")
}

func f0() {
	fmt.Println("run f0")
}
