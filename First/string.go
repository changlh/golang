package main

import (
	"fmt"
)

func reverse(str string) string {
	str1 := []rune(str)
	str2 := make([]rune, len(str1))
	for i := 0; i < len(str1); i++ {
		str2[i] = str1[len(str1)-i-1]
	}
	return string(str2)
}
func reverse1(str string) string {
	str1 := []rune(str)
	for i := 0; i < len(str1)/2; i++ {
		str1[i], str1[len(str1)-i-1] = str1[len(str1)-i-1], str1[i]
	}
	return string(str1)
}

func main() {
	s := "hello,world"
	fmt.Println(s)

	//只能替换成一个字符
	s1 := []byte(s)
	fmt.Println(s1)
	s1[1] = 'l'		//只能用单引号
	s = string(s1)
	fmt.Println(s)

	//可以替换为多个字符
	s2 := "hello,world"
	s = "hhh" + s2[1:]		//1:表示1到结尾
	fmt.Println(s)

	s3 := "hello,world"
	fmt.Println(reverse(s3))

	s4 := "你好，世界"
	fmt.Println(reverse1(s4))
}
