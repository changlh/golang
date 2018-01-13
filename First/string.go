package main

import "fmt"

func main() {
	s := "hello,world"
	fmt.Println(s)

	//只能替换成一个字符
	s1 := []byte(s)
	s1[1] = 'l'		//只能用单引号
	s = string(s1)
	fmt.Println(s)

	//可以替换为多个字符
	s2 := "hello,world"
	s = "hhh" + s2[1:]		//1:表示1到结尾 
	fmt.Println(s)
}
