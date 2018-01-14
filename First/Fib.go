package main

import "fmt"

func forFib() {
	a := 1
	b := 1
	for {
		a, b = b, a+b
		if a >= 100 {
			break
		}
		fmt.Println(a)
	}
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	forFib()
	fmt.Println("fib:",fib(10))
}
