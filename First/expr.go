package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//func expr(a, b int, ops string) (res int) {
//	if ops == "+" {
//		fmt.Println(a + b)
//	} else if ops == "-" {
//		fmt.Println(a - b)
//	} else if ops == "*" {
//		fmt.Println(a * b)
//	} else if ops == "/" {
//		fmt.Println(a / b)
//	}
//	return
//}

func main() {
	stra := os.Args[1]
	op := os.Args[2]
	strb := os.Args[3]
	var a, b int
	var err error
	if a, err = strconv.Atoi(stra); err != nil {
		log.Fatal(err)
	}
	if b, err = strconv.Atoi(strb); err != nil {
		log.Fatal(err)
	}
	//expr(a, b, op)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("请输入正确的运算符:", op)
	}
}
