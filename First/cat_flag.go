package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)
//此程序后续需要判断没加参数的情况

//useage信息
func useage() {
	fmt.Println("useage:\n" +
		"   COMMAND OPTIONS FILENAME")
	fmt.Println(" OPTIONS:\n" +
		"\t-file=\t filename")
}

//读取文件内容
func readfile(f string) {
	filecont, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(filecont))
}

func main() {
	//判断是否输入参数
	if len(os.Args) != 2 {
		useage()
		return
	}
	//增加参数
	filename := flag.String("file", "", "pls input filename")
	flag.Parse()
	fmt.Println("Filename is :", *filename)
	readfile(*filename)
	//fmt.Println("Args is :", flag.Args())
}

