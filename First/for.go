package main

import "fmt"

func main() {
	s := "hello"
	for i, a := range s {
		fmt.Println("i =", i, "a =", ";", a)
	}
	for _, b := range s {
		fmt.Println("b =", b)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	var c = true
	for c {
		c = false
	}
	for {
		break //或者continue，上面循环都能使用break和continue
	}
}
