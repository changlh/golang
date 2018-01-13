package main

import (
	"os"
	"strconv"
	"fmt"
	//"testing"
)

func andargs(x , y int)(z int) {
	z = x + y
	return
}

func main () {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	fmt.Println("a =",a)
	fmt.Println("b =",b)
	fmt.Println("a + b =",andargs(a, b))
}