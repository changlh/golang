package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"net/http"
)

func useage() {
	fmt.Println("useage:\n" +
		"   COMMAND FILENAME1 [ FILENAME2 ... ]")
	fmt.Println(" OPTIONS:\n" +
		"\tFILENAME\t 可以是文件，也可以是http url")
}

//打印文件内容
func Printfile(name string) {
	filename, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(filename))
}

//打印http_url内容
//对网页内容打印处理有问题
func Printhttp(url string) {
	urlcont, err := http.Get(url)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
	web, err := ioutil.ReadAll(urlcont.Body)
	if err != nil{
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(string(web))
	return
}

func main() {
	if len(os.Args) == 1 {
		useage()
		return
	}
	for i := 1; i < len(os.Args); i++ {
		filename := os.Args[i]
		if strings.HasPrefix(filename, "http://") {
			Printhttp(filename)
		} else {
			Printfile(filename)
		}
	}
}
