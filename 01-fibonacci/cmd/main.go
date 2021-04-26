package main

import (
	"fmt"
	fibonacci "thinknetica/01-fibonacci/pkg/fibonacci"
)

func main() {
	n := 21
	res := fibonacci.Calc(&n)
	fmt.Printf("Fib num %d - %d", n, res)
}