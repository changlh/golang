package main

import (
	"fmt"
	//"os"
	//"io/ioutil"
	"os"
	"io/ioutil"
)

//useage信息
func useage(){
	fmt.Println("useage:\n\t cat filename")
}
//读取文件内容
func readfile(filename string) {
	filecont , err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(filecont))
}

func main(){
	//判断是否输入参数
	if len(os.Args) != 2 {
		useage()
		return
	}
	a := os.Args[1]
	//fmt.Println(a)
	readfile(a)
}